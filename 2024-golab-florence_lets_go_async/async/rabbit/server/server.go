package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

type orderDetails struct{}

func main() {
	http.HandleFunc("/order", handleOrder)
	fmt.Println("Server listening on port 8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
func handleOrder(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Order processing started")

	details := getOrderDetails()
	err := generateInvoice(details)
	if err != nil {
		fmt.Println("Failed to send invoice generate message: " + err.Error())
	}
	handover(details)

	fmt.Println("Order processed")
	_, _ = fmt.Fprint(w, "Order processed")
}
func getOrderDetails() orderDetails {
	return orderDetails{}
}
func generateInvoice(details orderDetails) error {
	amqpConfig := amqp.NewDurableQueueConfig("amqp://guest:guest@localhost:5672/")
	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		return err
	}

	reqBodyBytes, _ := json.Marshal(details)
	msg := message.NewMessage(watermill.NewUUID(), reqBodyBytes)
	if err = publisher.Publish("printer_queue", msg); err != nil {
		return err
	}

	return nil
}
func handover(_ orderDetails) {
}
