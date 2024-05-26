package controllers

import (
	"backend/initializers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {
	var body struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Active      bool    `json:"active"`
		TeacherIDs  []uint  `json:"teacher_ids"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var teachers []models.Teacher
	initializers.DB.Where("id IN ?", body.TeacherIDs).Find(&teachers)

	course := models.Course{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Active:      body.Active,
		Teachers:    teachers,
	}

	if initializers.DB.Create(&course).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating the course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course created succesfully"})
}

func AssociateTeacherCourse(c *gin.Context) {
	var body struct {
		TeacherID uint `json:"teacher_id"`
		CursoID   uint `json:"curso_id"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var curso models.Course
	if initializers.DB.First(&curso, body.CursoID).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course not found"})
		return
	}

	var teacher models.Teacher
	if initializers.DB.First(&teacher, body.TeacherID).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not find teacher"})
		return
	}

	if initializers.DB.Model(&curso).Association("Teachers").Append(&teacher) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not associate the teacher to the course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher associated to the course successfully"})
}
