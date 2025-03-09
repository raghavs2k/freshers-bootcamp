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

var lastOrderTimes = make(map[string]time.Time) // Track last order per customer

// PlaceOrder - Create a new order
func PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.Mutex.Lock()
	defer utils.Mutex.Unlock()

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

	// Check stock availability
	if product.Quantity < order.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock available"})
		return
	}

	// Deduct stock
	product.Quantity -= order.Quantity
	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
		return
	}

	// Generate unique order ID
	order.ID = "ORD" + uuid.New().String()
	order.Status = "order placed"
	order.CreatedAt = time.Now()

	// Save order in DB
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Log transaction with Order ID
	transaction := models.Transaction{
		ID:         "TXN" + order.ID,
		OrderID:    order.ID, // Store Order ID in Transaction
		CustomerID: order.CustomerID,
		ProductID:  order.ProductID,
		Quantity:   order.Quantity,
		TotalPrice: float64(order.Quantity) * product.Price,
		Status:     "order placed",
		CreatedAt:  time.Now(),
	}
	config.DB.Create(&transaction)

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

// DeleteOrder - Deletes an order using its transaction record
func DeleteOrder(c *gin.Context) {
	utils.Mutex.Lock()
	defer utils.Mutex.Unlock()

	orderID := c.Param("id")

	// Find the order in transactions
	var transaction models.Transaction
	if err := config.DB.First(&transaction, "order_id = ?", orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction record not found for this order"})
		return
	}

	// Find the order before deleting
	var order models.Order
	if err := config.DB.First(&order, "id = ?", transaction.OrderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Delete the order
	if err := config.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	// Update transaction status to "order deleted"
	transaction.Status = "order deleted"
	if err := config.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
