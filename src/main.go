package main

import (
	"os"
	"net/http"
)

func sendText(w http.ResponseWriter, r *http.Request) {
	text, _ := os.Hostname()
	text= text + "\n"
	w.Write([]byte(text))
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}
	return ":" + port
}

func main() {
	http.HandleFunc("/", sendText)
	http.ListenAndServe(port(), nil)
}
