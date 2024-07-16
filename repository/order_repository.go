// repository/order_repository.go
package repository

import (
	"context"
	"order-management-system/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	DB *mongo.Database
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) CreateOrder(order *models.Order) error {
	collection := repo.DB.Collection("orders")
	order.OrderDate = time.Now()
	_, err := collection.InsertOne(context.Background(), order)
	return err
}

func (repo *OrderRepository) GetAllOrders() ([]models.Order, error) {
	collection := repo.DB.Collection("orders")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var orders []models.Order
	for cursor.Next(context.Background()) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (repo *OrderRepository) GetOrderById(orderID primitive.ObjectID) (*models.Order, error) {
	collection := repo.DB.Collection("orders")
	var order models.Order
	err := collection.FindOne(context.Background(), bson.M{"_id": orderID}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
