package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID        int
	orderDate time.Time
	customer  string
	amount    float64
}

func main() {
	order := Order{
		ID:        1,
		orderDate: time.Now(),
		customer:  "John Doe",
		amount:    123.45,
	}

	// Print the order details
	fmt.Printf("Order ID: %d\n", order.ID)
	fmt.Printf("Order Date: %s\n", order.orderDate)                                           // Print the time directly
	fmt.Printf("Order Date (formatted): %s\n", order.orderDate.Format("2006-01-02 15:04:05")) // Format the time
	fmt.Printf("Customer: %s\n", order.customer)
	fmt.Printf("Amount: %.2f\n", order.amount)
}
