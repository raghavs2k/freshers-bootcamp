package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Student Table
type Student struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	DOB       uint    `json:"dob"`
	Address   string  `json:"address"`
	Marks     []Marks `gorm:"foreignKey:StudentID"`
}

// Subject Table
type Subject struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

// Marks Table
type Marks struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	StudentID uint    `json:"student_id"`
	SubjectID uint    `json:"subject_id"`
	Marks     uint    `json:"marks"`
	Student   Student `gorm:"foreignKey:StudentID"`
	Subject   Subject `gorm:"foreignKey:SubjectID"`
}

func GetAllStudents(db *gorm.DB, students *[]Student) error {
	return db.Preload("Marks.Subject").Find(students).Error
}

func CreateStudent(db *gorm.DB, student *Student) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// First create the student
		if err := tx.Create(student).Error; err != nil {
			return err
		}

		// Then create marks
		for i := range student.Marks {
			// Verify subject exists
			var subject Subject
			if err := tx.First(&subject, student.Marks[i].SubjectID).Error; err != nil {
				return fmt.Errorf("subject with ID %d does not exist", student.Marks[i].SubjectID)
			}

			// Set the student ID for the mark and remove any ID to let GORM auto-increment
			student.Marks[i].StudentID = student.ID
			student.Marks[i].ID = 0 // Reset ID to let GORM auto-increment

			// Create the mark
			if err := tx.Create(&student.Marks[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
func GetStudentByID(db *gorm.DB, student *Student, id uint) error {
	return db.Preload("Marks.Subject").Where("id = ?", id).First(student).Error
}

func UpdateStudent(db *gorm.DB, student *Student) error {
	return db.Save(student).Error
}

func DeleteStudent(db *gorm.DB, id uint) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// Delete marks associated with the student
		if err := tx.Where("student_id = ?", id).Delete(&Marks{}).Error; err != nil {
			return err
		}

		// Delete the student
		if err := tx.Delete(&Student{}, id).Error; err != nil {
			return err
		}

		return nil
	})
}

func GetAllMarks(db *gorm.DB, marks *[]Marks) error {
	return db.Find(marks).Error
}

func GetAllSubjects(db *gorm.DB, subjects *[]Subject) error {
	return db.Find(subjects).Error
}
