.PHONY: build run

build:
	docker build -t go-rabbit-printer-demo .

run:
	docker rm -f rabbit-printer; docker run -e GOGC=off --memory=64m --network host --name rabbit-printer go-rabbit-printer-demo

RABBITMQ_DSN := "amqp://guest:guest@localhost:5672/"

rabbit: ## Starts the rabbitmq instance in docker container on port 5672, credentials are guest:guest
	docker rm -f rabbit; docker run --name rabbit --memory=2g -p 5672:5672 -p 15672:15672 -d rabbitmq:4.0.2-management