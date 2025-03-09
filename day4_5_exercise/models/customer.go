package models

import "gorm.io/gorm"

type Customer struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func MigrateCustomers(db *gorm.DB) {
	db.AutoMigrate(&Customer{})
}
