package controllers

import (
	"day4/config"
	"day4/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTransactionHistory - Fetch all orders as transactions
func GetTransactionHistory(c *gin.Context) {
	var orders []models.Order

	// Fetch all orders from DB
	if err := config.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	// Return transaction history
	c.JSON(http.StatusOK, gin.H{"transactions": orders})
}
