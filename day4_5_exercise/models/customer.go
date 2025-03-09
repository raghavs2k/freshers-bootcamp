package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func MigrateCustomers(db *gorm.DB) {
	db.AutoMigrate(&Customer{})
}
