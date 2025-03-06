package main

import (
	"io/ioutil"
	"net/http"
	"time"

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

	auth := gin.BasicAuth(gin.Accounts{
		"user":  "pass",
		"user1": "pass1",
	})

	router.GET("/getURLData/:name/:age", getURLData)

	admin := router.Group("/admin", auth)
	{
		admin.GET("/getData", getData)
	}
	client := router.Group("/client")
	{
		client.GET("/getQueryData", getQueryData)
	}

	server := &http.Server{
		Addr:         ":9091",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	} //server configuration using http and gin

	server.ListenAndServe()
}
