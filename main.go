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

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Database connection
	ipaddress := "192.168.103.162"

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + ipaddress + ":27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//database init
	db := client.Database("orderdb")

	//mqtt client init
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://" + ipaddress + ":1883")
	opts.SetClientID("go_mqtt_client")

	// Create and start an MQTT client
	mqttClient := MQTT.NewClient(opts)

	mqttService := service.NewMQTTService(mqttClient)

	// Initialize repositories and services
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, mqttService)
	orderService.ListenForOrderUpdates(ctx)

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
