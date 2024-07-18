// api/handler/order_handler.go
package handler

import (
	"fmt"
	"net/http"
	"order-management-system/models"
	"order-management-system/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderHandler struct {
	Service     *service.OrderService
	MQTTService *service.MQTTService
}

func NewOrderHandler(router *gin.Engine, service *service.OrderService) {
	handler := &OrderHandler{Service: service}

	router.POST("/orders", handler.CreateOrder)
	router.GET("/orders", handler.GetAllOrders)
	router.GET("/orders/:id", handler.GetOrderById)
	router.DELETE("/orders/:id", handler.DeleteOrder)
	router.PUT("/orders/:id", handler.UpdateOrder) // Update order
	router.POST("/orders/:id/process", handler.ProcessOrder)

}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.OrderID = primitive.NewObjectID()
	if err := h.Service.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.Service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetOrderById(c *gin.Context) {
	orderID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := h.Service.GetOrderById(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	orderID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	err = h.Service.DeleteOrder(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	orderID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.OrderID = orderID
	if err := h.Service.UpdateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

func (h *OrderHandler) ProcessOrder(c *gin.Context) {
	orderID, err := primitive.ObjectIDFromHex(c.Param("id"))
	//fmt.Println("........", orderID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	resultChan := h.Service.ProcessOrder(orderID)

	status := <-resultChan
	// fmt.Println(status)

	if status == "Order processed successfully" {
		c.JSON(http.StatusOK, gin.H{"status": status})
	} else {
		fmt.Print("error sending to result chan")
		c.JSON(http.StatusInternalServerError, gin.H{"status": status})
	}
}
