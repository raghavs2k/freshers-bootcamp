package routes

import (
	"day4/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	productRoutes := router.Group("/product")
	{
		productRoutes.POST("/", controllers.CreateProduct)
		productRoutes.PATCH("/:id", controllers.UpdateProduct)
		productRoutes.GET("/:id", controllers.GetProduct)
		productRoutes.DELETE("/:id", controllers.DeleteProduct)
	}

	router.GET("/products", controllers.GetAllProducts)
}
