package controllers

import (
	"day4/config"
	"day4/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTransactionHistory - Fetch all transactions (orders placed or deleted)
func GetTransactionHistory(c *gin.Context) {
	var transactions []models.Transaction

	// Fetch all transactions from DB
	if err := config.DB.Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	// Return transaction history
	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
