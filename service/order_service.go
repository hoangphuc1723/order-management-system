// service/order_service.go
package service

import (
	"context"
	"order-management-system/models"
	"order-management-system/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	Repo        *repository.OrderRepository
	MQTTService *MQTTService
}

func NewOrderService(repo *repository.OrderRepository, mqttService *MQTTService) *OrderService {
	return &OrderService{Repo: repo, MQTTService: mqttService}
}

func (s *OrderService) ProcessOrder(ctx context.Context, orderID primitive.ObjectID) chan string {
	result := make(chan string)

	go func() {
		defer close(result)

		// Update the order status in the database
		err := s.Repo.UpdateOrderStatus(ctx, orderID, "Processed")
		if err != nil {
			result <- "Failed to process order"
			return
		}

		// Publish MQTT message
		err = s.MQTTService.Publish("orders/processed", orderID.Hex())
		if err != nil {
			result <- "Failed to send MQTT message"
			return
		}

		result <- "Order processed successfully"
	}()

	return result
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
