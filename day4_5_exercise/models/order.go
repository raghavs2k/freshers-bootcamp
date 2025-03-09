package models

import (
	"time"

	"log"

	"gorm.io/gorm"
)

// Order struct (Parent Table)
type Order struct {
	ID         string      `json:"id" gorm:"primaryKey"`
	CustomerID string      `json:"customer_id"`
	TotalPrice float64     `json:"total_price"`
	Status     string      `json:"status"` // "order placed", "processed", "failed"
	CreatedAt  time.Time   `json:"created_at"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

// OrderItem struct (Child Table)
type OrderItem struct {
	ID        string  `json:"id" gorm:"primaryKey"`
	OrderID   string  `json:"order_id" gorm:"index"` // Foreign Key
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"` // Price at the time of order
}

// TableName - Specify custom table name for OrderItem
func (OrderItem) TableName() string {
	return "order_items"
}

// MigrateOrders - Run migration
func MigrateOrders(db *gorm.DB) {
	// AutoMigrate will only add missing fields/tables, won't delete/modify existing columns
	err := db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		log.Printf("Error migrating database: %v", err)
	}
}
