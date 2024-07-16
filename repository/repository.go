// repository/order_repository.go
package repository

import (
    "database/sql"
    "order-management/models"
)

type OrderRepository struct {
    DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
    return &OrderRepository{DB: db}
}

func (repo *OrderRepository) CreateOrder(order *models.Order) error {
    _, err := repo.DB.Exec("INSERT INTO orders (customer_id, order_date, status, total_amount, shipping_address_id, payment_id) VALUES (?, ?, ?, ?, ?, ?)",
        order.CustomerID, order.OrderDate, order.Status, order.TotalAmount, order.ShippingAddressID, order.PaymentID)
    return err
}

func (repo *OrderRepository) GetAllOrders() ([]models.Order, error) {
    rows, err := repo.DB.Query("SELECT * FROM orders")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    orders := []models.Order{}
    for rows.Next() {
        var order models.Order
        if err := rows.Scan(&order.OrderID, &order.CustomerID, &order.OrderDate, &order.Status, &order.TotalAmount, &order.ShippingAddressID, &order.PaymentID); err != nil {
            return nil, err
        }
        orders = append(orders, order)
    }
    return orders, nil
}

func (repo *OrderRepository) GetOrderById(orderID int) (*models.Order, error) {
    var order models.Order
    err := repo.DB.QueryRow("SELECT * FROM orders WHERE order_id = ?", orderID).Scan(&order.OrderID, &order.CustomerID, &order.OrderDate, &order.Status, &order.TotalAmount, &order.ShippingAddressID, &order.PaymentID)
    if err != nil {
        return nil, err
    }
    return &order, nil
}
