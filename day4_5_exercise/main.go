package main

import (
	"day4/config"
	"day4/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	routes.ProductRoutes(r)
	routes.CustomerRoutes(r)
	routes.OrderRoutes(r)
	routes.TransactionRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}

}
