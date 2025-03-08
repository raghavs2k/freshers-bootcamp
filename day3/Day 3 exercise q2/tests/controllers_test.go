package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"example.com/m/config"
	"example.com/m/controllers"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestGetStudentsController(t *testing.T) {
	assert := assert.New(t)
	setupTestDB(t)
	router := setupTestRouter()
	router.GET("/api/students", controllers.GetStudents)

	req, _ := http.NewRequest("GET", "/api/students", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code, "Expected status OK")

	var response []models.Student
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(err, "Failed to parse response")
}

func TestCreateStudentController(t *testing.T) {
	assert := assert.New(t)
	setupTestDB(t)
	router := setupTestRouter()
	router.POST("/api/students", controllers.CreateStudent)

	student := models.Student{
		FirstName: "John",
		LastName:  "Doe",
		DOB:       1626354614,
		Address:   "Test St",
		Marks: []models.Marks{
			{SubjectID: 1, Marks: 85},
			{SubjectID: 2, Marks: 90},
		},
	}

	body, err := json.Marshal(student)
	assert.NoError(err, "Failed to marshal student")

	req, _ := http.NewRequest("POST", "/api/students", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code, "Expected status OK")
}

func TestGetStudentByIDController(t *testing.T) {
	assert := assert.New(t)
	setupTestDB(t)
	router := setupTestRouter()
	router.GET("/api/students/:id", controllers.GetStudentByID)

	// Create test student
	student := models.Student{
		FirstName: "Test",
		LastName:  "User",
		DOB:       1626354614,
		Address:   "Test Address",
		Marks: []models.Marks{
			{SubjectID: 1, Marks: 85},
			{SubjectID: 2, Marks: 90},
		},
	}
	err := models.CreateStudent(config.DB, &student)
	assert.NoError(err, "Failed to create test student")

	req, _ := http.NewRequest("GET", "/api/students/"+strconv.Itoa(int(student.ID)), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code, "Expected status OK")
}

func TestUpdateStudentController(t *testing.T) {
	assert := assert.New(t)
	setupTestDB(t)
	router := setupTestRouter()
	router.PUT("/api/students/:id", controllers.UpdateStudent)

	// Create test student
	student := models.Student{
		FirstName: "Test",
		LastName:  "User",
		DOB:       1626354614,
		Address:   "Test Address",
		Marks: []models.Marks{
			{SubjectID: 1, Marks: 85},
			{SubjectID: 2, Marks: 90},
		},
	}
	err := models.CreateStudent(config.DB, &student)
	assert.NoError(err, "Failed to create test student")

	updateData := models.Student{
		FirstName: "Updated",
		LastName:  "User",
		DOB:       1626354614,
		Address:   "Updated Address",
		Marks: []models.Marks{
			{SubjectID: 1, Marks: 75},
			{SubjectID: 2, Marks: 80},
		},
	}

	body, err := json.Marshal(updateData)
	assert.NoError(err, "Failed to marshal update data")

	req, _ := http.NewRequest("PUT", "/api/students/"+strconv.Itoa(int(student.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code, "Expected status OK")
}

func TestDeleteStudentController(t *testing.T) {
	assert := assert.New(t)
	setupTestDB(t)
	router := setupTestRouter()
	router.DELETE("/api/students/:id", controllers.DeleteStudent)

	// Create test student
	student := models.Student{
		FirstName: "Test",
		LastName:  "User",
		DOB:       1626354614,
		Address:   "Test Address",
	}
	result := config.DB.Create(&student)
	assert.NoError(result.Error, "Failed to create test student")

	req, _ := http.NewRequest("DELETE", "/api/students/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code, "Expected status OK")
}

func TestGetMarksController(t *testing.T) {
	assert := assert.New(t)
	setupTestDB(t)
	router := setupTestRouter()
	router.GET("/api/marks", controllers.GetMarks)

	req, _ := http.NewRequest("GET", "/api/marks", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code, "Expected status OK")

	var response []models.Marks
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(err, "Failed to parse response")
}

func TestGetSubjectsController(t *testing.T) {
	assert := assert.New(t)
	setupTestDB(t)
	router := setupTestRouter()
	router.GET("/api/subjects", controllers.GetSubjects)

	req, _ := http.NewRequest("GET", "/api/subjects", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code, "Expected status OK")

	var response []models.Subject
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(err, "Failed to parse response")
}
