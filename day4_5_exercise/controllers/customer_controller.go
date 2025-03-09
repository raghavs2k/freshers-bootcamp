package controllers

import (
	"day4/config"
	"day4/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateCustomer - Add a new customer
func CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique customer ID
	customer.ID = "CST" + uuid.New().String()

	// Save customer to DB
	if err := config.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

// GetCustomer - Retrieve a customer by ID
func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer

	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// GetAllCustomers - Retrieve all customers
func GetAllCustomers(c *gin.Context) {
	var customers []models.Customer
	config.DB.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"customers": customers})
}
