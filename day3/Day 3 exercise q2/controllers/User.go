package controllers

import (
	"net/http"
	"strconv"

	"example.com/m/config"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	if err := models.GetAllStudents(config.DB, &students); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Students not found"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := models.CreateStudent(config.DB, &student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func GetStudentByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var student models.Student
	if err := models.GetStudentByID(config.DB, &student, uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	student.ID = uint(id)

	if err := models.UpdateStudent(config.DB, &student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := models.DeleteStudent(config.DB, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}

func GetMarks(c *gin.Context) {
	var marks []models.Marks
	if err := models.GetAllMarks(config.DB, &marks); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Marks not found"})
		return
	}
	c.JSON(http.StatusOK, marks)
}

func GetSubjects(c *gin.Context) {
	var subjects []models.Subject
	if err := models.GetAllSubjects(config.DB, &subjects); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subjects not found"})
		return
	}
	c.JSON(http.StatusOK, subjects)
}
