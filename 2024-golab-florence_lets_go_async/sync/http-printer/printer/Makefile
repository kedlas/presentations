.PHONY: build run

build:
	docker build -t go-http-printer-printer-demo .

run:
	docker run -p 8091:8091 -e GOGC=off --memory=128m --network host --name http-printer-printer go-http-printer-printer-demo
