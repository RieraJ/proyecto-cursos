package controllers

import (
	"backend/initializers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTeacher(c *gin.Context) {
	var body struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Crear un nuevo profesor
	newTeacher := models.Teacher{
		Name:    body.Name,
		Surname: body.Surname,
	}

	if initializers.DB.Create(&newTeacher).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed while creating a teacher"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher created successfully"})
}

func GetAllTeachers(c *gin.Context) {
	var teachers []models.Teacher
	initializers.DB.Find(&teachers)

	c.JSON(http.StatusOK, teachers)
}

func GetTeacherByID(c *gin.Context) {
	var teacher models.Teacher
	if initializers.DB.First(&teacher, c.Param("id")).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}

	c.JSON(http.StatusOK, teacher)
}
