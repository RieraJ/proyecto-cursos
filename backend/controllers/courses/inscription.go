package controllers

import (
	"backend/dto"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnrollUser(c *gin.Context) {
	var enrollUser dto.InscriptionRequest
	err := c.BindJSON(&enrollUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.InscriptionServiceInterfaceInstance.EnrollUser(enrollUser)
	if err != nil {
		if err.Error() == "user is already enrolled in this course" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User enrolled successfully"})
}
