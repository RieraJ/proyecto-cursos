package users

import (
	"backend/dto"
	service "backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	// Get the email/pass of request body
	var body dto.SignUpRequest
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to bind JSON: %s", err.Error()),
		})
		return
	}

	// Call the service
	err = service.UserServiceInterfaceInstance.Signup(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to signup: %s", err.Error())})
		return
	}

	// Return the response
	result := dto.SignUpResponse{Message: "User created successfully"}
	c.JSON(http.StatusOK, result)
}
