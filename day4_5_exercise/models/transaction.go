package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	OrderID    string    `json:"order_id"`
	CustomerID string    `json:"customer_id"`
	ProductID  string    `json:"product_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

func MigrateTransaction(db *gorm.DB) {
	db.AutoMigrate(&Transaction{})
}
