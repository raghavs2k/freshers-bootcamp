package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am GIN Framework",
	})
}

func getPostData(c *gin.Context) {
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)
	c.JSON(200, gin.H{
		"bodyData": string(value),
	})
}
func getQueryData(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func getURLData(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	router := gin.Default()

	router.GET("/getData", getData)
	router.POST("/getPostData", getPostData)
	router.GET("/getQueryData", getQueryData)
	router.GET("/getURLData/:name/:age", getURLData)

	router.Run()
}
