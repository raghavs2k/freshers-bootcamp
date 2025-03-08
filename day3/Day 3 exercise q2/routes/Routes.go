package routes

import (
	"example.com/m/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.GET("/students", controllers.GetStudents)
		grp1.POST("/students", controllers.CreateStudent)
		grp1.GET("/students/:id", controllers.GetStudentByID)
		grp1.PUT("/students/:id", controllers.UpdateStudent)
		grp1.DELETE("/students/:id", controllers.DeleteStudent)
		grp1.GET("/marks", controllers.GetMarks)
		grp1.GET("/subjects", controllers.GetSubjects)
	}
	return r
}
