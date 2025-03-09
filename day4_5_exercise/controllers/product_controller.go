package controllers

import (
	"day4/config"
	"day4/models"
	"day4/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateProduct - Add a new product
func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique product ID
	product.ID = "PROD" + uuid.New().String()

	// Save product to DB
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":           product.ID,
		"product_name": product.ProductName,
		"price":        product.Price,
		"quantity":     product.Quantity,
		"message":      "Product successfully added",
	})
}

// UpdateProduct - Modify price and quantity
func UpdateProduct(c *gin.Context) {
	utils.Mutex.Lock()         // Lock before update
	defer utils.Mutex.Unlock() // Unlock after update

	var product models.Product
	id := c.Param("id")

	// Find product
	if err := config.DB.First(&product, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Bind JSON
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Save updated product
	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

// GetProduct - Retrieve a product by ID
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.First(&product, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetAllProducts - Retrieve all products
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// Delete Product (DELETE /product/:id)
func DeleteProduct(c *gin.Context) {
	utils.Mutex.Lock()         // Lock before delete
	defer utils.Mutex.Unlock() // Unlock after delete

	id := c.Param("id")
	if err := config.DB.Delete(&models.Product{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
