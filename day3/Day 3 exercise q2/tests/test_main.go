package tests

import (
	"testing"

	"example.com/m/config"
	"example.com/m/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) {
	// Load test environment variables
	if err := godotenv.Load("../.env.test"); err != nil {
		t.Fatalf("Error loading .env.test file: %v", err)
	}

	// Initialize test database connection
	dbConfig := config.BuildDBConfig()
	dsn := config.DbURL(dbConfig)

	var err error
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Drop tables in correct order
	config.DB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	config.DB.Exec("DROP TABLE IF EXISTS marks")
	config.DB.Exec("DROP TABLE IF EXISTS students")
	config.DB.Exec("DROP TABLE IF EXISTS subjects")
	config.DB.Exec("SET FOREIGN_KEY_CHECKS = 1")

	// Migrate tables
	config.DB.AutoMigrate(&models.Student{}, &models.Subject{}, &models.Marks{})

	// Insert test subjects
	subjects := []models.Subject{
		{ID: 1, Name: "Mathematics"},
		{ID: 2, Name: "Science"},
		{ID: 3, Name: "English"},
	}
	for _, subject := range subjects {
		if err := config.DB.Create(&subject).Error; err != nil {
			t.Fatalf("Failed to create test subject: %v", err)
		}
	}
}
