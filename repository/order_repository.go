// repository/order_repository.go
package repository

import (
	"context"
	"order-management-system/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"fmt"
)

type OrderRepository struct {
	DB *mongo.Database
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	repo := &OrderRepository{DB: db}
	repo.InitializeOrders()
	return repo
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

		fmt.Printf("Retrieved order: %+v\n", order)
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

func (r *OrderRepository) DeleteOrder(orderID primitive.ObjectID) error {
	collection := r.DB.Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": orderID})
	return err
}

// UpdateOrderStatus updates the status of an order
func (r *OrderRepository) UpdateOrderStatus(ctx context.Context, orderID primitive.ObjectID, status string) error {
	filter := bson.M{"_id": orderID}
	update := bson.M{"$set": bson.M{"status": status}}
	_, err := r.DB.Collection("orders").UpdateOne(ctx, filter, update)
	return err
}

func (r *OrderRepository) InitializeOrders() error {
	collection := r.DB.Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if there are already orders in the collection
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	if count == 0 {
		orders := []models.Order{
			{
				CustomerID:        "1",
				OrderDate:         time.Now(),
				Status:            "Pending",
				TotalAmount:       100.0,
				ShippingAddressID: "1",
				PaymentID:         "1",
			},
			{
				CustomerID:        "2",
				OrderDate:         time.Now(),
				Status:            "Completed",
				TotalAmount:       150.0,
				ShippingAddressID: "2",
				PaymentID:         "2",
			},
			{
				CustomerID:        "3",
				OrderDate:         time.Now(),
				Status:            "Shipped",
				TotalAmount:       200.0,
				ShippingAddressID: "3",
				PaymentID:         "3",
			},
			{
				CustomerID:        "4",
				OrderDate:         time.Now(),
				Status:            "Processing",
				TotalAmount:       250.0,
				ShippingAddressID: "4",
				PaymentID:         "4",
			},
			{
				CustomerID:        "5",
				OrderDate:         time.Now(),
				Status:            "Cancelled",
				TotalAmount:       300.0,
				ShippingAddressID: "5",
				PaymentID:         "5",
			},
			{
				CustomerID:        "6",
				OrderDate:         time.Now(),
				Status:            "Returned",
				TotalAmount:       350.0,
				ShippingAddressID: "6",
				PaymentID:         "6",
			},
			{
				CustomerID:        "7",
				OrderDate:         time.Now(),
				Status:            "Pending",
				TotalAmount:       400.0,
				ShippingAddressID: "7",
				PaymentID:         "7",
			},
		}

		for _, order := range orders {
			_, err := collection.InsertOne(ctx, order)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *OrderRepository) UpdateOrder(order models.Order) error {
	collection := r.DB.Collection("orders")
	filter := bson.M{"_id": order.OrderID}
	update := bson.M{
		"$set": order,
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}
