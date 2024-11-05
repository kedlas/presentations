package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
	go createAndSendInvoice(details)
	handover(details)

	fmt.Println("Order processed")
	_, _ = fmt.Fprint(w, "Order processed")
}
func getOrderDetails() orderDetails {
	return orderDetails{}
}
func createAndSendInvoice(details orderDetails) {
	reqBodyBytes := new(bytes.Buffer)
	_ = json.NewEncoder(reqBodyBytes).Encode(details)

	_, err := http.Post(
		"http://127.0.0.1:8091/print",
		"application/json",
		reqBodyBytes,
	)
	if err != nil {
		fmt.Println("Failed to generate invoice")
	}
}

func handover(_ orderDetails) {
}
