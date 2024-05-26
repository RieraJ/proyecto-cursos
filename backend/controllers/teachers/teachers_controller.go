package controllers

import (
	"backend/initializers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTeacher(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Crear un nuevo profesor
	newTeacher := models.Teacher{
		Name: body.Name,
	}

	if initializers.DB.Create(&newTeacher).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed while creating a teacher"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher created successfully"})
}
