// models/order.go
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID           primitive.ObjectID `bson:"_id,omitempty" json:"order_id"`
	CustomerID        primitive.ObjectID `bson:"customer_id,omitempty" json:"customer_id"`
	OrderDate         time.Time          `bson:"order_date" json:"order_date"`
	Status            string             `bson:"status" json:"status"`
	TotalAmount       float64            `bson:"total_amount" json:"total_amount"`
	ShippingAddressID primitive.ObjectID `bson:"shipping_address_id,omitempty" json:"shipping_address_id"`
	PaymentID         primitive.ObjectID `bson:"payment_id,omitempty" json:"payment_id"`
}
