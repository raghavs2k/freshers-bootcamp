package routes

import (
	"day4/controllers"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(router *gin.Engine) {
	customerRoutes := router.Group("/customer")
	{
		customerRoutes.POST("/", controllers.CreateCustomer)
		customerRoutes.GET("/:id", controllers.GetCustomer)
		customerRoutes.PATCH("/:id", controllers.UpdateCustomer)
		customerRoutes.DELETE("/:id", controllers.DeleteCustomer)
	}

	router.GET("/customers", controllers.GetAllCustomers)
}
