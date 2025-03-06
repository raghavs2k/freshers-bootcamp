package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sharma-raghav/golang-bootcamp/day3/gin_framework/video_3/middleware"
)

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am GIN Framework getData",
	})
}
func getData1(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am GIN Framework getData1",
	})
}
func getData2(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am GIN Framework getData2",
	})
}

func main() {
	router := gin.New() //router without logging and default middleware

	// router.Use(middleware.Authenticate) //Applying to every route

	admin := router.Group("/admin", middleware.Authenticate(), middleware.AddHeader)
	{
		admin.GET("/getData1", getData1)
		admin.GET("/getData2", getData2)
	}

	router.GET("/getData", middleware.Authenticate(), getData)
	router.Run(":8081")
}
