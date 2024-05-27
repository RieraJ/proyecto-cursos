package controllers

import (
	"backend/initializers"
	"backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func EnrollStudent(c *gin.Context) {
	var body struct {
		UserID   uint `json:"user_id"`
		CourseID uint `json:"course_id"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Verificar si el usuario y el curso existen
	var user models.User
	var course models.Course

	if initializers.DB.First(&user, body.UserID).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if initializers.DB.First(&course, body.CourseID).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	// Crear la inscripción
	inscription := models.CourseInscription{
		UserID:          body.UserID,
		CourseID:        body.CourseID,
		InscriptionDate: time.Now(),
	}

	if initializers.DB.Create(&inscription).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enroll student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student enrolled successfully"})
}
