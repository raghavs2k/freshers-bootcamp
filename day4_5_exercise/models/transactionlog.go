package models

import "gorm.io/gorm"

type TransactionLog struct {
	ID        string  `gorm:"primaryKey" json:"id"`
	OrderID   string  `json:"order_id"`
	ProductID string  `json:"product_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
}

func MigrateTransactions(db *gorm.DB) {
	db.AutoMigrate(&TransactionLog{})
}
