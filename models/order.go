// models/order.go
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID           primitive.ObjectID `bson:"_id,omitempty" json:"orderId"`
	CustomerID        string             `bson:"customerId,omitempty" json:"customerId"`
	OrderDate         time.Time          `bson:"orderDate" json:"orderDate"`
	Status            string             `bson:"status" json:"status"`
	TotalAmount       float64            `bson:"totalAmount" json:"totalAmount"`
	ShippingAddressID string             `bson:"shippingAddressId,omitempty" json:"shippingAddressId"`
	PaymentID         string             `bson:"paymentId,omitempty" json:"paymentId"`
}
