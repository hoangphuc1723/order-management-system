// service/order_service.go
package service

import (
    "order-management/models"
    "order-management/repository"
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

func (s *OrderService) GetOrderById(orderID int) (*models.Order, error) {
    return s.Repo.GetOrderById(orderID)
}
