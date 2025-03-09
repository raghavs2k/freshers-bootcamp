package routes

import (
	"day4/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/order")
	{
		orderRoutes.POST("/", controllers.PlaceOrder)
		orderRoutes.GET("/:id", controllers.GetOrder)
		orderRoutes.DELETE("/:id", controllers.DeleteOrder)

	}
}
