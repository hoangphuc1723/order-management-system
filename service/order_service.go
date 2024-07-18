// service/order_service.go
package service

import (
	"context"
	"fmt"
	"order-management-system/models"
	"order-management-system/repository"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	Repo        *repository.OrderRepository
	MQTTService *MQTTService
	mu          sync.Mutex
}

func NewOrderService(repo *repository.OrderRepository, mqttService *MQTTService) *OrderService {
	mqttService.Connect()
	mqttService.Subscribe("orders/processed", 2, mqttService.OnMessageReceived)

	return &OrderService{Repo: repo, MQTTService: mqttService}
}

func (s *OrderService) ProcessOrder(orderID primitive.ObjectID) chan string {
	result := make(chan string)

	go func() {
		defer func() {
			fmt.Println("MQTT", s.MQTTService.LastMessage)
			close(result)
		}()

		// Publish MQTT message
		s.MQTTService.Publish("orders/processed", 2, false, orderID.Hex())

		result <- "Order processed successfully"
	}()

	return result
}

func (s *OrderService) ListenForOrderUpdates(ctx context.Context) {
	go func() {
		for {
			select {
			case msg := <-s.MQTTService.MessageChan:
				fmt.Printf("Received MQTT message: %s\n", msg)
				orderID, err := primitive.ObjectIDFromHex(msg)
				if err == nil {
					// Update order status in the database
					err := s.Repo.UpdateOrderStatus(orderID, "Processed")
					if err != nil {
						fmt.Printf("Failed to update order: %s\n", err)
					} else {
						fmt.Printf("Order %s updated successfully\n", orderID.Hex())
					}
				}
			}
		}
	}()
}

func (s *OrderService) UpdateOrderStatus(orderID primitive.ObjectID, status string) error {
	return s.Repo.UpdateOrderStatus(orderID, status)
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.Repo.CreateOrder(order)
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.Repo.GetAllOrders()
}

func (s *OrderService) GetOrderById(orderID primitive.ObjectID) (*models.Order, error) {
	return s.Repo.GetOrderById(orderID)
}

func (s *OrderService) DeleteOrder(orderID primitive.ObjectID) error {
	return s.Repo.DeleteOrder(orderID)
}

func (s *OrderService) UpdateOrder(order models.Order) error {
	return s.Repo.UpdateOrder(order)
}
