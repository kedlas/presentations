package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.dataddo.com/pgq"
	"go.dataddo.com/pgq/x/schema"
	"net/http"
)

type orderDetails struct{}

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

	http.HandleFunc("/order", handleOrder)
	fmt.Println("Server listening on port 8090")
	err = http.ListenAndServe(":8090", nil)
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
	db, err := sql.Open("pgx", "postgresql://pgq:pgq@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	publisher := pgq.NewPublisher(db)

	reqBodyBytes, _ := json.Marshal(details)
	msg := &pgq.MessageOutgoing{Payload: json.RawMessage(reqBodyBytes)}
	_, err = publisher.Publish(context.Background(), "printer_queue", msg)
	if err != nil {
		return err
	}

	return nil
}
func handover(_ orderDetails) {
}
