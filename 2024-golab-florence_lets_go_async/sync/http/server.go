package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/order", handleOrder)
	_ = http.ListenAndServe(":8090", nil)
}
func handleOrder(w http.ResponseWriter, _ *http.Request) {
	details := getOrderDetails()
	generateInvoice(details)
	handover(details)

	_, _ = fmt.Fprint(w, "Order processed")
}
func getOrderDetails() orderDetails {
	time.Sleep(1 * time.Second)

	return orderDetails{}
}
func handover(_ orderDetails) {
	time.Sleep(1 * time.Second)
}
func generateInvoice(_ orderDetails) invoice {
	time.Sleep(5 * time.Second)

	return invoice{}
}

type orderDetails struct{}
type invoice struct{}
