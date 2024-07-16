// api/handler/order_handler.go
package handler

import (
    "net/http"
    "order-management/models"
    "order-management/service"
    "github.com/gin-gonic/gin"
)

type OrderHandler struct {
    Service *service.OrderService
}

func NewOrderHandler(router *gin.Engine, service *service.OrderService) {
    handler := &OrderHandler{Service: service}

    router.POST("/orders", handler.CreateOrder)
    router.GET("/orders", handler.GetAllOrders)
    router.GET("/orders/:id", handler.GetOrderById)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

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
    orderID := c.Param("id")

    order, err := h.Service.GetOrderById(orderID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, order)
}
