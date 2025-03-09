package controllers

import (
	"day4/config"
	"day4/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransactionHistory(c *gin.Context) {
	var transactions []models.Transaction

	if err := config.DB.Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
