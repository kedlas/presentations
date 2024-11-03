.PHONY: build run

build:
	docker build -t go-http-printer-server-demo .

run:
	docker rm -f http-printer-server; docker run -p 8090:8090 -e GOGC=off --memory=128m --network host --name http-printer-server go-http-printer-server-demo

create-order:
	curl --location --request POST 'http://localhost:8090/order'

create-multiple-orders:
	ab -n 10 -c 5 -m "POST" 'http://localhost:8090/order'