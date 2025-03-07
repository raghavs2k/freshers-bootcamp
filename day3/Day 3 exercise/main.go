package main

import (
	"fmt"

	"example.com/m/config"
	"example.com/m/models"
	"example.com/m/routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	dsn := config.DbURL(config.BuildDBConfig())
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer func() {
		db, _ := config.DB.DB()
		db.Close()
	}()
	config.DB.AutoMigrate(&models.User{})
	r := routes.SetupRouter()
	r.Run()
}
