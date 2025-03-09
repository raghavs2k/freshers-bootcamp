package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	CustomerID string    `json:"customer_id"`
	ProductID  string    `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Status     string    `json:"status"` // "order placed", "processed", "failed"
	CreatedAt  time.Time `json:"created_at"`
}

// MigrateOrders - Run migration
func MigrateOrders(db *gorm.DB) {
	db.AutoMigrate(&Order{})
}
