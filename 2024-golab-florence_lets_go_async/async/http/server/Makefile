.PHONY: build run

build:
	docker build -t go-http-async-printer-server-demo .

run:
	docker rm -f http-async-printer-server; docker run -p 8090:8090 --network host --name http-async-printer-server go-http-async-printer-server-demo

create-order:
	curl --location --request POST 'http://localhost:8090/order'

create-multiple-orders:
	ab -n 10 -c 5 -m "POST" 'http://localhost:8090/order'