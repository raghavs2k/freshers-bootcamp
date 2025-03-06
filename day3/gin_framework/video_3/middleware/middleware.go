package middleware

import "github.com/gin-gonic/gin"

func Authenticate() gin.HandlerFunc {
	//write custom logic to be applied before my middleware is executed
	return func(c *gin.Context) {
		if !(c.Request.Header.Get("Token") == "Auth") {
			c.AbortWithStatusJSON(500, gin.H{
				"Message": "Token not found",
			})
			return
		}
		c.Next()

	}
}

func AddHeader(c *gin.Context) {
	c.Writer.Header().Set("Key", "Value")
	c.Next()
}

// func Authenticate(c *gin.Context) {
// 	if !(c.Request.Header.Get("Token") == "Auth") {
// 		c.AbortWithStatusJSON(500, gin.H{
// 			"Message": "Token not found",
// 		})
// 		return
// 	}
// 	c.Next()
// }
