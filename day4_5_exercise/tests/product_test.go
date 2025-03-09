package tests

import (
	"day4/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	mockDB := new(MockDB)
	product := models.Product{
		ID:          "PROD67890",
		ProductName: "Basket",
		Price:       100,
		Quantity:    30,
	}

	mockDB.On("Create", mock.Anything).Return(&gorm.DB{})

	result := mockDB.Create(&product)

	assert.NotNil(t, result)
	mockDB.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	mockDB := new(MockDB)
	product := models.Product{
		ID:          "PROD67890",
		ProductName: "Updated Basket",
		Price:       60,
		Quantity:    50,
	}

	mockDB.On("Save", mock.Anything).Return(&gorm.DB{})

	result := mockDB.Save(&product)

	assert.NotNil(t, result)
	mockDB.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	mockDB := new(MockDB)
	product := models.Product{ID: "PROD67890"}

	mockDB.On("Delete", mock.Anything, mock.Anything).Return(&gorm.DB{})

	result := mockDB.Delete(&product, "id = ?", "PROD67890")

	assert.NotNil(t, result)
	mockDB.AssertExpectations(t)
}
