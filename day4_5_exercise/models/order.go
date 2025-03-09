package models

import (
	"time"

	"log"

	"gorm.io/gorm"
)

type Order struct {
	ID         string      `json:"id" gorm:"primaryKey"`
	CustomerID string      `json:"customer_id"`
	TotalPrice float64     `json:"total_price"`
	Status     string      `json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        string  `json:"id" gorm:"primaryKey"`
	OrderID   string  `json:"order_id" gorm:"index"`
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

func (OrderItem) TableName() string {
	return "order_items"
}

func MigrateOrders(db *gorm.DB) {

	err := db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		log.Printf("Error migrating database: %v", err)
	}
}
