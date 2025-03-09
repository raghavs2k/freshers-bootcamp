package controllers

import (
	"day4/config"
	"day4/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTransactionHistory - Fetch all transactions (orders with items)
func GetTransactionHistory(c *gin.Context) {
	var orders []models.Order

	// Change "OrderItems" to "Items" to match the struct field name
	if err := config.DB.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	// Return transaction history
	c.JSON(http.StatusOK, gin.H{"transactions": orders})
}
