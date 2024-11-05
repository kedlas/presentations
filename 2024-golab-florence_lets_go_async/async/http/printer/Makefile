.PHONY: build run logs

build:
	docker build -t go-http-async-printer-printer-demo .

run:
	docker rm -f http-async-printer-printer; docker run -e GOGC=off --memory=64m --network host --name http-async-printer-printer go-http-async-printer-printer-demo

logs:
	docker logs http-async-printer-printer