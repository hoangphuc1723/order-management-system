// service/order_processing.go
package service

import (
	"log"
	"order-management-system/models"
	"sync"
)

var orderChannel = make(chan models.Order, 100)
var orderMutex = &sync.Mutex{}

func (s *OrderService) StartOrderProcessing() {
	go func() {
		for order := range orderChannel {
			processOrder(order)
		}
	}()
}

func processOrder(order models.Order) {
	orderMutex.Lock()
	defer orderMutex.Unlock()

	// Simulate order processing
	log.Printf("Processing order ID: %s\n", order.OrderID.Hex())
	// Update order status
}
