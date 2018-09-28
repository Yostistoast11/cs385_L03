all: app

app: 
	docker build -t app .
	docker run -d --name goapp app

clean:
	docker stop goapp
	docker rm goapp

prune:
	docker system prune -a
