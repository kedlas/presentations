package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
	err, inv := createInvoice(details)
	if err != nil {
		fmt.Println("Failed to generate invoice")
		_, _ = fmt.Fprint(w, "Order process failed: failed to generate invoice")
		return
	}
	handover(details, inv)

	fmt.Println("Order processed")
	_, _ = fmt.Fprint(w, "Order processed")
}
func getOrderDetails() orderDetails {
	return orderDetails{}
}
func createInvoice(details orderDetails) (error, invoice) {
	reqBodyBytes := new(bytes.Buffer)
	_ = json.NewEncoder(reqBodyBytes).Encode(details)

	resp, err := http.Post(
		"http://127.0.0.1:8091/print",
		"application/json",
		reqBodyBytes,
	)
	if err != nil {
		return err, invoice{}
	}

	return nil, parsePrinterResponse(resp)
}
func handover(_ orderDetails, _ invoice) {
}

func parsePrinterResponse(_ *http.Response) invoice {
	return invoice{}
}
