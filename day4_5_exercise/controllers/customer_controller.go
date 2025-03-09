package controllers

import (
	"day4/config"
	"day4/models"
	"day4/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.ID = "CST" + uuid.New().String()

	if err := config.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer

	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func GetAllCustomers(c *gin.Context) {
	var customers []models.Customer
	config.DB.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

func UpdateCustomer(c *gin.Context) {
	utils.Mutex.Lock()
	defer utils.Mutex.Unlock()

	var customer models.Customer
	id := c.Param("id")

	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	config.DB.Save(&customer)
	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	utils.Mutex.Lock()
	defer utils.Mutex.Unlock()

	id := c.Param("id")
	if err := config.DB.Delete(&models.Customer{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
