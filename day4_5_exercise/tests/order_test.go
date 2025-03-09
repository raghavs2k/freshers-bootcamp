package tests

import (
	"day4/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateOrder(t *testing.T) {
	mockDB := new(MockDB)
	order := models.Order{
		ID:         "ORD12345",
		CustomerID: "CST12345",
		Status:     "order placed",
	}

	mockDB.On("Create", mock.Anything).Return(&gorm.DB{})

	result := mockDB.Create(&order)

	assert.NotNil(t, result)
	mockDB.AssertExpectations(t)
}

func TestDeleteOrder(t *testing.T) {
	mockDB := new(MockDB)
	order := models.Order{ID: "ORD12345"}

	mockDB.On("Delete", mock.Anything, mock.Anything).Return(&gorm.DB{})

	result := mockDB.Delete(&order, "id = ?", "ORD12345")

	assert.NotNil(t, result)
	mockDB.AssertExpectations(t)
}
