// models/order.go
package models

import "time"

type Order struct {
    OrderID           int       `json:"order_id"`
    CustomerID        int       `json:"customer_id"`
    OrderDate         time.Time `json:"order_date"`
    Status            string    `json:"status"`
    TotalAmount       float64   `json:"total_amount"`
    ShippingAddressID int       `json:"shipping_address_id"`
    PaymentID         int       `json:"payment_id"`
}
