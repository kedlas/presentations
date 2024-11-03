.PHONY: build run logs

build:
	docker build -t go-http-printer-printer-demo .

run:
	docker rm -f http-printer-printer; docker run -e GOGC=off --memory=64m --network host --name http-printer-printer go-http-printer-printer-demo

logs:
	docker logs http-printer-printer