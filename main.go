// main.go
package main

import (
	"context"
	"log"
	"net/http"
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

	mqttService := service.NewMQTTService("tcp://broker.hivemq.com:1883", "orderManagementClient")

	db := client.Database("orderdb")

	// Initialize repositories and services
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, mqttService)

	// Initialize Gin router
	r := gin.Default()

	// Set up routes
	handler.NewOrderHandler(r, orderService)

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	})

	// Serve static files (HTML, CSS, JS, images, etc.) from the "web" directory
	r.Static("/web", "./web")

	// Run the server
	r.Run(":8080")

	// Other initializations...
}
