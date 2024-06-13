package controllers

import (
	"backend/dto"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EnrollUser(c *gin.Context) {
	var request dto.InscriptionRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Obtener el userID de la cookie
	userIdStr, err := c.Cookie("userId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID cookie not found"})
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	request.UserID = uint(userId)

	if request.UserID == 0 || request.CourseID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID and CourseID are required"})
		return
	}

	if err := services.InscriptionServiceInterfaceInstance.EnrollUser(request); err != nil {
		if err.Error() == "user is already enrolled in this course" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enroll user: " + err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User enrolled successfully"})
}
