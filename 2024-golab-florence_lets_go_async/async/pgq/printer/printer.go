package main

import (
	"context"
	"database/sql"
	"fmt"
	"go.dataddo.com/pgq"
	"go.dataddo.com/pgq/x/schema"
	"runtime"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "postgresql://pgq:pgq@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	ctx := context.Background()
	_, err = db.ExecContext(ctx, schema.GenerateCreateTableQuery("printer_queue"))
	if err != nil {
		panic(err.Error())
	}

	h := handler{}
	consumer, err := pgq.NewConsumer(db, "printer_queue", &h,
		pgq.WithMaxParallelMessages(1),
		pgq.WithLockDuration(1*time.Minute),
		pgq.WithPollingInterval(500*time.Millisecond),
		// add other options here if you wish, please see the docs https://github.com/dataddo/pgq#consumer-options
	)
	if err != nil {
		panic(err.Error())
	}

	err = consumer.Run(ctx)
	if err != nil {
		panic(err.Error())
	}
}

type handler struct{}

func (h *handler) HandleMessage(_ context.Context, msg *pgq.MessageIncoming) (processed bool, err error) {
	fmt.Println("print message received")
	err, details := parseMessage(msg)
	if err != nil {
		return pgq.MessageNotProcessed, err
	}

	err = handlePrint(details)
	if err != nil {
		return pgq.MessageProcessed, err
	}

	return pgq.MessageProcessed, nil
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
func parseMessage(_ *pgq.MessageIncoming) (error, orderDetails) {
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
