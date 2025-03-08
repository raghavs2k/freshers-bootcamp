package main

import (
	"fmt"

	"example.com/m/config"
	"example.com/m/models"
	"example.com/m/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := config.DbURL(config.BuildDBConfig())
	var err error
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Database connection failed:", err)
	}

	config.DB.AutoMigrate(&models.Student{}, &models.Subject{}, &models.Marks{})

	// Insert default subjects if they don't exist
	subjects := []models.Subject{
		{ID: 1, Name: "Mathematics"},
		{ID: 2, Name: "Science"},
		{ID: 3, Name: "English"},
		{ID: 4, Name: "Social Science"},
		{ID: 5, Name: "Hindi"},
		// Add more subjects as needed
	}

	for _, subject := range subjects {
		config.DB.FirstOrCreate(&subject, models.Subject{ID: subject.ID})
	}
}

func main() {
	InitDB()

	r := routes.SetupRouter()
	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}
