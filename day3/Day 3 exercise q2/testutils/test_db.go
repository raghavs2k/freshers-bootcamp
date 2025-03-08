package testutils

import (
	"log"

	"example.com/m/config"
	"example.com/m/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func init() {
	if err := godotenv.Load("../.env.test"); err != nil {
		log.Printf("Error loading .env.test file: %v", err)
	}

	dbConfig := config.BuildDBConfig()
	dsn := config.DbURL(dbConfig)

	var err error
	TestDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Test database connection failed: %v", err)
	}

	TestDB.AutoMigrate(&models.Student{}, &models.Subject{}, &models.Marks{})
	config.DB = TestDB
}
