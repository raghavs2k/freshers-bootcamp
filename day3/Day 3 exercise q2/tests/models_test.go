package tests

import (
	"testing"

	"example.com/m/models"
	"example.com/m/testutils"
)

func TestCreateStudent(t *testing.T) {
	student := models.Student{
		FirstName: "John",
		LastName:  "Doe",
		DOB:       1626354614,
		Address:   "123 Test Street",
		Marks: []models.Marks{
			{SubjectID: 1, Marks: 85},
			{SubjectID: 2, Marks: 90},
		},
	}

	err := models.CreateStudent(testutils.TestDB, &student)
	got := err
	want := error(nil)

	if got != want {
		t.Errorf("TestCreateStudent failed: got %v, want %v", got, want)
	}

	if student.ID == 0 {
		t.Errorf("TestCreateStudent failed: got ID = 0, expected a valid ID")
	}
}

func TestGetAllStudents(t *testing.T) {
	var students []models.Student
	err := models.GetAllStudents(testutils.TestDB, &students)

	got := err
	want := error(nil)

	if got != want {
		t.Errorf("TestGetAllStudents failed: got %v, want %v", got, want)
	}

	if len(students) < 0 {
		t.Errorf("TestGetAllStudents failed: got length %d, expected at least 0", len(students))
	}
}

func TestGetStudentByID(t *testing.T) {
	student := models.Student{
		FirstName: "Alice",
		LastName:  "Smith",
		DOB:       1626354614,
		Address:   "456 Test Avenue",
		Marks: []models.Marks{
			{SubjectID: 1, Marks: 75},
			{SubjectID: 2, Marks: 80},
		},
	}
	_ = models.CreateStudent(testutils.TestDB, &student)

	var fetchedStudent models.Student
	err := models.GetStudentByID(testutils.TestDB, &fetchedStudent, student.ID)

	got := err
	want := error(nil)

	if got != want {
		t.Errorf("TestGetStudentByID failed: got %v, want %v", got, want)
	}

	if fetchedStudent.FirstName != student.FirstName {
		t.Errorf("TestGetStudentByID failed: got FirstName %s, want %s", fetchedStudent.FirstName, student.FirstName)
	}
}

func TestUpdateStudent(t *testing.T) {
	student := models.Student{
		FirstName: "Bob",
		LastName:  "Brown",
		DOB:       1626354614,
		Address:   "789 Test Road",
		Marks: []models.Marks{
			{SubjectID: 1, Marks: 65},
			{SubjectID: 2, Marks: 70},
		},
	}
	_ = models.CreateStudent(testutils.TestDB, &student)

	updatedStudent := models.Student{
		ID:        student.ID,
		FirstName: "Bobby",
		LastName:  "Brown",
		DOB:       student.DOB,
		Address:   "Updated Address",
		Marks: []models.Marks{
			{SubjectID: 1, Marks: 75},
			{SubjectID: 2, Marks: 80},
		},
	}
	err := models.UpdateStudent(testutils.TestDB, &updatedStudent)

	got := err
	want := error(nil)

	if got != want {
		t.Errorf("TestUpdateStudent failed: got %v, want %v", got, want)
	}

	var fetchedStudent models.Student
	_ = models.GetStudentByID(testutils.TestDB, &fetchedStudent, student.ID)

	if fetchedStudent.FirstName != "Bobby" {
		t.Errorf("TestUpdateStudent failed: got FirstName %s, want Bobby", fetchedStudent.FirstName)
	}
}

func TestDeleteStudent(t *testing.T) {
	student := models.Student{
		FirstName: "Charlie",
		LastName:  "Johnson",
		DOB:       1626354614,
		Address:   "101 Test Blvd",
		Marks: []models.Marks{
			{SubjectID: 1, Marks: 55},
			{SubjectID: 2, Marks: 60},
		},
	}
	_ = models.CreateStudent(testutils.TestDB, &student)

	err := models.DeleteStudent(testutils.TestDB, student.ID)

	got := err
	want := error(nil)

	if got != want {
		t.Errorf("TestDeleteStudent failed: got %v, want %v", got, want)
	}

	var fetchedStudent models.Student
	err = models.GetStudentByID(testutils.TestDB, &fetchedStudent, student.ID)

	if err == nil {
		t.Errorf("TestDeleteStudent failed: student still exists after deletion")
	}
}

func TestGetAllMarks(t *testing.T) {
	var marks []models.Marks
	err := models.GetAllMarks(testutils.TestDB, &marks)

	got := err
	want := error(nil)

	if got != want {
		t.Errorf("TestGetAllMarks failed: got %v, want %v", got, want)
	}

	if len(marks) < 0 {
		t.Errorf("TestGetAllMarks failed: got length %d, expected at least 0", len(marks))
	}
}

func TestGetAllSubjects(t *testing.T) {
	var subjects []models.Subject
	err := models.GetAllSubjects(testutils.TestDB, &subjects)

	got := err
	want := error(nil)

	if got != want {
		t.Errorf("TestGetAllSubjects failed: got %v, want %v", got, want)
	}

	if len(subjects) < 0 {
		t.Errorf("TestGetAllSubjects failed: got length %d, expected at least 0", len(subjects))
	}
}
