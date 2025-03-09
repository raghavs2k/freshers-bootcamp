package controllers

import (
	"day4/config"
	"day4/models"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var orderLock sync.Mutex
var lastOrderTimes = make(map[string]time.Time) // Track last order per customer

// PlaceOrder - Create a new order
func PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderLock.Lock()
	defer orderLock.Unlock()

	// Check cool-down period
	lastOrderTime, exists := lastOrderTimes[order.CustomerID]
	if exists && time.Since(lastOrderTime) < 5*time.Minute {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Please wait 5 minutes before placing another order"})
		return
	}

	// Fetch product details
	var product models.Product
	if err := config.DB.First(&product, "id = ?", order.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Check if enough stock is available
	if product.Quantity < order.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock available"})
		return
	}

	// Deduct stock & save
	product.Quantity -= order.Quantity
	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
		return
	}

	// Create Order
	order.ID = "ORD" + uuid.New().String()
	order.Status = "order placed"
	order.CreatedAt = time.Now()
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Update last order time
	lastOrderTimes[order.CustomerID] = time.Now()

	c.JSON(http.StatusCreated, gin.H{
		"id":         order.ID,
		"product_id": order.ProductID,
		"quantity":   order.Quantity,
		"status":     order.Status,
	})
}

// GetOrder - Retrieve a single order
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := config.DB.First(&order, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
