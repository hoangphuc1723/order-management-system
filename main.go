// main.go
package main

import (
    "database/sql"
    "log"
    "order-management/api/handler"
    "order-management/repository"
    "order-management/service"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Other initializations...

    // Start order processing
    orderService.StartOrderProcessing()
    
    // Database connection
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/orderdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Initialize repositories and services
    orderRepo := repository.NewOrderRepository(db)
    orderService := service.NewOrderService(orderRepo)

    // Initialize Gin router
    r := gin.Default()

    // Set up routes
    handler.NewOrderHandler(r, orderService)

    // Run the server
    r.Run(":8080")
}
