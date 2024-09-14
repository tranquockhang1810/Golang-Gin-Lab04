// Create controllers

package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tranquockhang1810/Golang-Gin-Lab04/db"
	"github.com/tranquockhang1810/Golang-Gin-Lab04/models"
)

//Get all students
func GetStudents(c *gin.Context) {
	students := db.Students
	c.JSON(http.StatusOK, gin.H{
		"data": students,
		"message": "Get students successfully",
	})
}

//Get student by id
func GetStudentById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"data": db.Students[id],
		"message": "Get student successfully",
	})
}

//Create student
func CreateStudent(c *gin.Context) {
	var newStudent models.Student

	// Parse JSON body vào struct Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid JSON data",
		})
		return
	}

	// Gán ID mới cho sinh viên
	newStudent.Id = len(db.Students) + 1

	// Lưu vào danh sách students
	db.Students[strconv.Itoa(newStudent.Id)] = newStudent

	c.JSON(http.StatusOK, gin.H{
		"data":    newStudent,
		"message": "Create student successfully",
	})
}

//Update student
func UpdateStudent(c *gin.Context) { 
	var student models.Student

	id := c.Param("id")

	// Parse JSON body với struct Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid JSON data",
		})
		return
	}

	if _, ok := db.Students[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}
	student.Id, _ = strconv.Atoi(id)
	db.Students[id] = student

	c.JSON(http.StatusOK, gin.H{
		"data":    student,
		"message": "Update student successfully",
	})
}

//Delete student
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	if _, ok := db.Students[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}

	delete(db.Students, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete student successfully",
	})
}