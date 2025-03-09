package main

import (
	"day4/config"
	"day4/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	// Register Routes
	routes.ProductRoutes(r)

	// Start Server
	r.Run(":8080")
}
