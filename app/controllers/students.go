package controllers

import (
	"gin-framework/app/db"
	"gin-framework/app/models"
	"log"
	"net/http"

	ws "gin-framework/app/websocket"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new student
// @Description Create a new student
// @Tags students
// @Accept json
// @Produce json
// @Param student body models.StudentInput true "Student object that needs to be added"
// @Success 201 {object} models.ResponseSuccessCreate
// @Failure 400 {object} models.ResponseInvalid
// @Failure 500 {object} models.ResponseFailedCreate
// @Router /students [post]
func CreateStudent(c *gin.Context) {
	var studentRequest models.StudentInput
	var student models.Student

	if err := c.ShouldBindJSON(&studentRequest); err != nil {
		ws.SendWebSocketUpdate("Invalid request!")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request!",
		})
		return
	}

	student.Nim = studentRequest.Nim
	student.Name = studentRequest.Name
	student.Jurusan = studentRequest.Jurusan

	if err := db.DB.Create(&student).Error; err != nil {
		ws.SendWebSocketUpdate("Failed to create student!")
		log.Printf("Failed to create student!")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create student!",
		})
		return
	}

	ws.SendWebSocketUpdate("Success create student")
	log.Printf("Student created!")

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created student!",
	})
}

// @Summary Get student by ID
// @Description Get student details by ID
// @Tags students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} models.StudentResponseByID
// @Failure 204 {object} models.ResponseStudentNotFound
// @Router /students/{id} [get]
func GetStudentByID(c *gin.Context) {
	var student models.Student

	if err := db.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		ws.SendWebSocketUpdate("Student not found!")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found!",
		})
		return
	}

	ws.SendWebSocketUpdate("Success get student by ID")

	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}

// @Summary Get all students
// @Description Get a list of all students
// @Tags students
// @Accept json
// @Produce json
// @Success 200 {object} models.AllStudentResponse
// @Failure 204 {object} models.ResponseStudentNotFound
// @Router /students [get]
func ReadAllStudents(c *gin.Context) {
	var students []models.Student

	if err := db.DB.Find(&students).Error; err != nil {
		ws.SendWebSocketUpdate("Student not found!")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Students not found!",
		})
		return
	}

	ws.SendWebSocketUpdate("Success all students")

	c.JSON(http.StatusOK, gin.H{
		"students": students,
	})
}

// @Summary Update student by ID
// @Description Update student details by ID with JSON data
// @Tags students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param student body models.StudentUpdate true "Updated student object"
// @Success 200 {object} models.ResponseSuccessUpdate
// @Failure 400 {object} models.ResponseInvalid
// @Failure 204 {object} models.ResponseStudentNotFound
// @Failure 500 {object} models.ResponseFailedUpdate
// @Router /students/{id} [put]
func UpdateStudentByID(c *gin.Context) {
	var student models.Student
	var studentRequest models.StudentUpdate

	if err := db.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		ws.SendWebSocketUpdate("Student not found!")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found!",
		})
		return
	}

	if err := c.ShouldBindJSON(&studentRequest); err != nil {
		ws.SendWebSocketUpdate("Invalid request!")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request!",
		})
		return
	}

	student.Name = studentRequest.Name

	if err := db.DB.Save(&student).Error; err != nil {
		ws.SendWebSocketUpdate("Failed to update student!")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update student!",
		})
		return
	}

	ws.SendWebSocketUpdate("Success get all student")

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully update student!",
	})
}

// @Summary Delete student by ID
// @Description Delete student by ID
// @Tags students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} models.ResponseSuccessDelete
// @Failure 204 {object} models.ResponseStudentNotFound
// @Failure 500 {object} models.ResponseFailedDelete
// @Router /students/{id} [delete]
func DeleteStudentByID(c *gin.Context) {
	var student models.Student

	if err := db.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		ws.SendWebSocketUpdate("Student not found!")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found!",
		})
		return
	}

	if err := db.DB.Delete(&student).Error; err != nil {
		ws.SendWebSocketUpdate("Failed to delete student!")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete student!",
		})
		return
	}

	ws.SendWebSocketUpdate("Success delete student")

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete student!",
	})
}
