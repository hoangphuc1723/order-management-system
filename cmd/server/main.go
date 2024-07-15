package main

import (
	"log"

	"github.com/hoangphuc1723/order-management-system/internal/api"
	"github.com/hoangphuc1723/order-management-system/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Connect to MongoDB
	err := utils.InitMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Setup API endpoints
	api.SetupRoutes(router)

	// Run the HTTP server
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
