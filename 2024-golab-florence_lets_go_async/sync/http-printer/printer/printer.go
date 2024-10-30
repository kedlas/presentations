package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	http.HandleFunc("/print", handlePrint)
	fmt.Println("Printer listening on port 8091")
	_ = http.ListenAndServe(":8091", nil)
}
func handlePrint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("print request received")
	details := parseRequest(req)
	generated := generateInvoice(details)
	fmt.Println("invoice generated")

	_, _ = fmt.Fprint(w, generated)
}
func parseRequest(_ *http.Request) orderDetails {
	return orderDetails{}
}
func generateInvoice(_ orderDetails) invoice {
	memLeak()
	printMemUsage()
	return invoice{}
}

func memLeak() {
	for i := 0; i < 1000000; i++ {
		memGarbage := make([]byte, 0)
		memGarbage = append(memGarbage, []byte("Hello there, I am leaking.")...)
	}
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

type orderDetails struct{}
type invoice struct{}
