// service/order_service.go
package service

import (
	"context"
	"order-management-system/models"
	"order-management-system/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	Repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{Repo: repo}
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

func (s *OrderService) ProcessOrder(ctx context.Context, orderID primitive.ObjectID) <-chan string {
	result := make(chan string)

	go func() {
		defer close(result)

		// Simulate long-running task
		time.Sleep(10 * time.Second)

		// Update the order status to "Processed"
		err := s.Repo.UpdateOrderStatus(ctx, orderID, "Processed")
		if err != nil {
			result <- "Error processing order: " + err.Error()
			return
		}

		result <- "Order processed successfully"
	}()

	return result
}
