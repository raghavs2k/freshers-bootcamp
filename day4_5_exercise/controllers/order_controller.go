package controllers

import (
	"day4/config"
	"day4/models"
	"day4/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var lastOrderTimes = make(map[string]time.Time)

func PlaceOrder(c *gin.Context) {
	var req struct {
		CustomerID string `json:"customer_id"`
		Items      []struct {
			ProductID string `json:"product_id"`
			Quantity  int    `json:"quantity"`
		} `json:"items"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.Mutex.Lock()
	defer utils.Mutex.Unlock()

	// Cooldown Check
	lastOrderTime, exists := lastOrderTimes[req.CustomerID]
	if exists && time.Since(lastOrderTime) < 5*time.Minute {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Wait 5 minutes before placing another order"})
		return
	}

	var total float64
	var orderItems []models.OrderItem

	for _, item := range req.Items {
		var product models.Product
		if err := config.DB.First(&product, "id = ?", item.ProductID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		if product.Quantity < item.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock for product " + product.ID})
			return
		}

		product.Quantity -= item.Quantity
		config.DB.Save(&product)

		orderItems = append(orderItems, models.OrderItem{
			ID:        "ITEM" + uuid.New().String(),
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})

		total += product.Price * float64(item.Quantity)
	}

	order := models.Order{
		ID:         "ORD" + uuid.New().String(),
		CustomerID: req.CustomerID,
		TotalPrice: total,
		Status:     "order placed",
		CreatedAt:  time.Now(),
		Items:      orderItems,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	transaction := models.Transaction{
		ID:         "TXN" + uuid.New().String(),
		OrderID:    order.ID,
		CustomerID: order.CustomerID,
		TotalPrice: total,
		Status:     "order placed",
		CreatedAt:  time.Now(),
	}

	if err := config.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	for i := range orderItems {
		orderItems[i].OrderID = order.ID
	}
	config.DB.Create(&orderItems)

	lastOrderTimes[req.CustomerID] = time.Now()

	c.JSON(http.StatusCreated, gin.H{
		"order_id":    order.ID,
		"total_price": total,
		"status":      order.Status,
		"items":       orderItems,
	})
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := config.DB.Preload("Items").First(&order, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	utils.Mutex.Lock()
	defer utils.Mutex.Unlock()

	orderID := c.Param("id")

	var order models.Order
	if err := config.DB.First(&order, "id = ?", orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := config.DB.Where("order_id = ?", orderID).Delete(&models.OrderItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order items"})
		return
	}

	if err := config.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	transaction := models.Transaction{
		ID:         "TXN" + uuid.New().String(),
		OrderID:    orderID,
		CustomerID: order.CustomerID,
		TotalPrice: order.TotalPrice,
		Status:     "order deleted",
		CreatedAt:  time.Now(),
	}

	if err := config.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
