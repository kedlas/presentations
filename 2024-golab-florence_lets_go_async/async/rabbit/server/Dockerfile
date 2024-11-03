# syntax=docker/dockerfile:1

FROM golang:1.23

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY * ./

RUN go build -o /go-rabbit-server-demo

EXPOSE 8090

CMD [ "/go-rabbit-server-demo" ]