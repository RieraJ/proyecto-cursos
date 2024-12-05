package users

import (
	"backend/clients"
	"backend/dao"
	"backend/dto"
	service "backend/services"
	"fmt"
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
	token, err := service.UserServiceInterfaceInstance.Login(*client)
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

func UpdateUserType(c *gin.Context) {
	var requestBody struct {
		UserID   uint   `json:"user_id"`
		UserType string `json:"user_type"`
	}

	// Obtener el usuario autenticado desde el contexto
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Verificar si el usuario es admin
	currentUser := user.(dao.User)
	if currentUser.UserType != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Vincular el cuerpo de la solicitud al struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Llamar al servicio para actualizar el userType
	err := service.UserServiceInterfaceInstance.UpdateUserType(requestBody.UserID, requestBody.UserType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "User type updated successfully"})
}

func GetAllUsers(c *gin.Context) {
	users, err := service.UserServiceInterfaceInstance.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get users: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, users)
}
