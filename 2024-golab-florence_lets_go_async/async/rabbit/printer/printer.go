package main

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"log"
	"runtime"
	"time"
)

func main() {
	amqpConfig := amqp.NewDurableQueueConfig("amqp://guest:guest@localhost:5672/")
	subscriber, err := amqp.NewSubscriber(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		log.Fatalf("could not create subscriber: %v", err)
	}

	messages, err := subscriber.Subscribe(context.Background(), "printer_queue")
	if err != nil {
		log.Fatalf("could not subscribe to printer_queue: %v", err)
	}

	for msg := range messages {
		fmt.Println("print message received")
		err, details := parseMessage(msg)
		if err != nil {
			// ideally we should reject, byt watermill does not support that
			msg.Ack()
			continue
		}

		err = handlePrint(details)
		if err != nil {
			msg.Nack()
			continue
		}

		msg.Ack()
		fmt.Println("print message processed")
	}
}
func handlePrint(details orderDetails) error {
	err, inv := printInvoice(details)
	if err != nil {
		return err
	}
	err = sendInvoice(details, inv)
	if err != nil {
		return err
	}

	return nil
}
func parseMessage(_ *message.Message) (error, orderDetails) {
	return nil, orderDetails{}
}
func printInvoice(_ orderDetails) (error, invoice) {
	printMemUsage()
	memLeak()
	printMemUsage()

	runtime.GC()

	return nil, invoice{}
}
func sendInvoice(_ orderDetails, _ invoice) error {
	return nil
}

func memLeak() {
	for i := 0; i < 1000000; i++ {
		memGarbage := make([]byte, 0)
		memGarbage = append(memGarbage, []byte("Hello there Gophers, I am leaking a lot of memory here...")...)
	}
	time.Sleep(2 * time.Second)
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
