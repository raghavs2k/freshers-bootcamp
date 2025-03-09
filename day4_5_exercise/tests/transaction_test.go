package tests

import (
	"day4/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestGetTransactions(t *testing.T) {
	mockDB := new(MockDB)
	transactions := []models.Transaction{
		{ID: "TXN123", OrderID: "ORD12345", CustomerID: "CST12345", TotalPrice: 500},
	}

	mockDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{})

	result := mockDB.Find(&transactions)

	assert.NotNil(t, result)
	mockDB.AssertExpectations(t)
}
