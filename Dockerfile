FROM golang:1.9

COPY src/main.go src/

WORKDIR src/

CMD ["go", "run", "main.go"]
