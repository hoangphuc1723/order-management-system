// main.go
package main

import (
	"context"
	"log"
	"order-management-system/api/handler"
	"order-management-system/repository"
	"order-management-system/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Database connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("orderdb")

	// Initialize repositories and services
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)

	// Initialize Gin router
	r := gin.Default()

	// Set up routes
	handler.NewOrderHandler(r, orderService)

	// Run the server
	r.Run(":8080")

	// Other initializations...

	// Start order processing
	orderService.StartOrderProcessing()

	// Other code...

}
