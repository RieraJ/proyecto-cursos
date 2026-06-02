package controllers

import (
	"backend/dao"
	"backend/dto"
	"backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnrollUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	currentUser := user.(dao.User)

	var request dto.InscriptionRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	request.UserID = currentUser.ID

	if request.CourseID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CourseID is required"})
		return
	}

	if err := services.UserServiceInterfaceInstance.EnrollUser(request); err != nil {
		if err.Error() == "user is already enrolled in this course" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			log.Printf("EnrollUser error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enroll user"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User enrolled successfully"})
}
