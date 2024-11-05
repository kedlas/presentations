package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	http.HandleFunc("/print", handlePrint)
	fmt.Println("Printer listening on port 8091")
	_ = http.ListenAndServe(":8091", nil)
}
func handlePrint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("print request received")
	details := parseRequest(req)
	inv := printInvoice(details)
	fmt.Println("invoice generated")

	_, _ = fmt.Fprint(w, inv)
}
func parseRequest(_ *http.Request) orderDetails {
	return orderDetails{}
}
func printInvoice(_ orderDetails) invoice {
	printMemUsage()
	memLeak()
	printMemUsage()
	return invoice{}
}

func memLeak() {
	for i := 0; i < 1000000; i++ {
		memGarbage := make([]byte, 0)
		memGarbage = append(memGarbage, []byte("Hello there Gophers, I am leaking a lot of memory here...")...)
	}
	time.Sleep(2 * time.Second)
	runtime.GC()
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
