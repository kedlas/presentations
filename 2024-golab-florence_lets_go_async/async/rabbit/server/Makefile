.PHONY: build run rabbit

build:
	docker build -t go-rabbit-server-demo .

run:
	docker rm -f rabbit-server; docker run -p 8090:8090 --network host --name rabbit-server go-rabbit-server-demo

create-order:
	curl --location --request POST 'http://localhost:8090/order'

create-multiple-orders:
	ab -n 10 -c 5 -m "POST" 'http://localhost:8090/order'