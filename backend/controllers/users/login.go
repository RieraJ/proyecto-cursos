package users

import (
	"backend/clients"
	"backend/dto"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var client *dto.User
	// Bind the request body to the LoginRequest struct
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the Login service
	token, err := services.AuthServiceInterfaceInstance.Login(*client)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userDAO, _ := clients.SelectUserByEmail(client.Email) // Obtener el usuario para obtener el userId

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", token, 60*60*24*30, "/", "localhost", false, true)
	c.SetCookie("Auth", token, 60*60*24*30, "", "", false, true)
	c.SetCookie("userId", strconv.Itoa(int(userDAO.ID)), 60*60*24*30, "/", "localhost", false, true) // Guardar userId en cookie
	c.SetCookie("userId", strconv.Itoa(int(userDAO.ID)), 60*60*24*30, "/", "localhost", false, true) // Guardar userId en cookie
	c.JSON(http.StatusOK, gin.H{"message": "Cookie successfully generated"})
}
