package routes

import (
	"day4/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(router *gin.Engine) {
	router.GET("/transactions", controllers.GetTransactionHistory)
}
