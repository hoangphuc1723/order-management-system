package models

import (
    "fmt"
    "time"
)

type Order struct {
    orderID int
    orderDate time.Time
    customerID string
    Amount float64
}

func main() {
    order := Order{
        ID:        1,
        orderDate: time.Now(),
        Customer:  "John Doe",
        Amount:    123.45,
    }

    // Print the order details
    fmt.Printf("Order ID: %d\n", order.ID)
    fmt.Printf("Order Date: %s\n", order.orderDate.Format("2006-01-02 15:04:05"))
    fmt.Printf("Customer: %s\n", order.Customer)
    fmt.Printf("Amount: %.2f\n", order.Amount)
}
