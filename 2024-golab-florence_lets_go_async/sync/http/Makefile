.PHONY: build run

build:
	docker build -t go-http-demo .

run:
	docker run -p 8090:8090 -e GOGC=off --memory=128m  go-http-demo

create-order:
	curl --location --request POST 'http://localhost:8090/order'

create-multiple-orders:
	ab -n 10 -c 5 -m "POST" 'http://localhost:8090/order'