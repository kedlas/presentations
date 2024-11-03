.PHONY: help postgres rabbit

help: ## Print help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

POSTGRES_DSN := "postgresql://pgq:pgq@localhost:5432/postgres?sslmode=disable"
RABBITMQ_DSN := "amqp://guest:guest@localhost:5672/"

postgres: ## Starts the postgres 16 instance in docker container on port 5432
	docker rm -f pgq-postgres; docker run --name pgq-postgres --memory=2g -e POSTGRES_USER=pgq -e POSTGRES_PASSWORD=pgq -p 5432:5432 -d postgres:16.0

rabbit: ## Starts the rabbitmq instance in docker container on port 5672, credentials are guest:guest
	docker rm -f rabbit; docker run --name rabbit --memory=2g -p 5672:5672 -p 15672:15672 -d rabbitmq:4.0.2-management
