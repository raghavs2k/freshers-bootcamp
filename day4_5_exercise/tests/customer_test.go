package tests

import (
	"day4/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateCustomer(t *testing.T) {
	mockDB := new(MockDB)
	customer := models.Customer{
		ID:    "CST12345",
		Name:  "Rahul Sharma",
		Email: "rahul@gmail.com",
	}

	mockDB.On("Create", mock.Anything).Return(&gorm.DB{})

	result := mockDB.Create(&customer)

	assert.NotNil(t, result)
	mockDB.AssertExpectations(t)
}

func TestUpdateCustomer(t *testing.T) {
	mockDB := new(MockDB)
	customer := models.Customer{
		ID:    "CST12345",
		Name:  "Raghav Sharma",
		Email: "raghav2@gmail.com",
	}

	mockDB.On("Save", mock.Anything).Return(&gorm.DB{})

	result := mockDB.Save(&customer)

	assert.NotNil(t, result)
	mockDB.AssertExpectations(t)
}

func TestDeleteCustomer(t *testing.T) {
	mockDB := new(MockDB)
	customer := models.Customer{ID: "CST12345"}

	mockDB.On("Delete", mock.Anything, mock.Anything).Return(&gorm.DB{})

	result := mockDB.Delete(&customer, "id = ?", "CST12345")

	assert.NotNil(t, result)
	mockDB.AssertExpectations(t)
}
