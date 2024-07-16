package models

import (
	"time"
)

type Order struct {
	OrderID           string    `json:"orderId" bson:"orderId"`
	CustomerID        string    `json:"customerId" bson:"customerId"`
	OrderDate         time.Time `json:"orderDate" bson:"orderDate"`
	Status            string    `json:"status" bson:"status"`
	TotalAmount       float64   `json:"totalAmount" bson:"totalAmount"`
	ShippingAddressID string    `json:"shippingAddressId,omitempty" bson:"shippingAddressId,omitempty"`
	PaymentID         string    `json:"paymentId,omitempty" bson:"paymentId,omitempty"`
}
