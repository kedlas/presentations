package main

import (
	"fmt"
	"net/http"
	"time"
)

type orderDetails struct{}
type invoice struct{}

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
	inv := createInvoice(details)
	handover(details, inv)

	fmt.Println("Order processed")
	_, _ = fmt.Fprint(w, "Order processed")
}
func getOrderDetails() orderDetails {
	time.Sleep(100 * time.Millisecond)
	return orderDetails{}
}
func createInvoice(_ orderDetails) invoice {
	time.Sleep(5 * time.Second)
	return invoice{}
}
func handover(_ orderDetails, _ invoice) {
	time.Sleep(200 * time.Millisecond)
}
