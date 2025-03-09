package models

import "gorm.io/gorm"

type Product struct {
	ID          string  `gorm:"primaryKey" json:"id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func MigrateProducts(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}
