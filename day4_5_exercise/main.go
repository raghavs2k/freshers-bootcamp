package main

import (
	"day4/config"
	"day4/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	// Registering all Routes
	routes.ProductRoutes(r)
	routes.CustomerRoutes(r)
	routes.OrderRoutes(r)
	routes.TransactionRoutes(r)

	// Start Server
	r.Run(":8080")
}
