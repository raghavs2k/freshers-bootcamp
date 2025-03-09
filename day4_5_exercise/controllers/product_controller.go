package controllers

import (
	"day4/config"
	"day4/models"
	"day4/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = "PROD" + uuid.New().String()

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

func UpdateProduct(c *gin.Context) {
	utils.Mutex.Lock()
	defer utils.Mutex.Unlock()

	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.First(&product, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func DeleteProduct(c *gin.Context) {
	utils.Mutex.Lock()
	defer utils.Mutex.Unlock()

	id := c.Param("id")
	if err := config.DB.Delete(&models.Product{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
