package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/order", handleOrder)
	fmt.Println("Server listening on port 8090")
	_ = http.ListenAndServe(":8090", nil)
}
func handleOrder(w http.ResponseWriter, _ *http.Request) {
	details := getOrderDetails()
	generateInvoice(details)
	handover(details)

	_, _ = fmt.Fprint(w, "Order processed")
}
func getOrderDetails() orderDetails {
	return orderDetails{}
}
func handover(_ orderDetails) {
}
func generateInvoice(_ orderDetails) invoice {
	resp, err := http.Post(
		"http://127.0.0.1:8091/print",
		"application/json",
		nil,
	)
	if err != nil {
		panic(err)
	}

	return parsePrinterResponse(resp)
}

func parsePrinterResponse(_ *http.Response) invoice {
	return invoice{}
}

type orderDetails struct{}
type invoice struct{}
