package users

import (
	"backend/dto"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var body dto.LoginRequest
	// Bind the request body to the LoginRequest struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the Login service
	result, err := services.AuthServiceInterfaceInstance.Login(body)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set the token in a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", result.Token, 60*60*24*30, "", "", false, true)

	// Return the token in the response
	c.JSON(http.StatusOK, gin.H{"message": "Cookie succesfully generated"})

}
