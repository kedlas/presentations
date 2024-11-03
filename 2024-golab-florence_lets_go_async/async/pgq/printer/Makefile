.PHONY: build run rabbit

build:
	docker build -t go-pgq-printer-demo .

run:
	docker rm -f pgq-printer; docker run -p 8090:8090 -e GOGC=off --memory=128m --network host --name pgq-printer go-pgq-printer-demo

create-order:
	curl --location --request POST 'http://localhost:8090/order'

create-multiple-orders:
	ab -n 10 -c 5 -m "POST" 'http://localhost:8090/order'

POSTGRES_DSN := "postgresql://pgq:pgq@localhost:5432/postgres?sslmode=disable"

postgres: ## Starts the postgres 16 instance in docker container on port 5432
	docker rm -f pgq-postgres; docker run --name pgq-postgres --memory=2g -e POSTGRES_USER=pgq -e POSTGRES_PASSWORD=pgq -p 5432:5432 -d postgres:16.0