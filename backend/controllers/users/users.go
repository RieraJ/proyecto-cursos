package users

import (
	"backend/clients"
	"backend/dao"
	"backend/dto"
	service "backend/services"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const maxPhotoBytes = 5 << 20 // 5 MB

func Login(c *gin.Context) {
	var client *dto.User
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := service.UserServiceInterfaceInstance.Login(*client)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	userDAO, _ := clients.SelectUserByEmail(client.Email)

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", token, 60*60*24*30, "/", "", false, true)
	c.SetCookie("userId", strconv.Itoa(int(userDAO.ID)), 60*60*24*30, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Cookie successfully generated"})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.SetCookie("userId", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func Signup(c *gin.Context) {
	var body dto.SignUpRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := service.UserServiceInterfaceInstance.Signup(body); err != nil {
		if err.Error() == "user already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}
		// Validation errors from service are safe to forward
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SignUpResponse{Message: "User created successfully"})
}

func UpdateUserType(c *gin.Context) {
	var requestBody struct {
		UserID   uint   `json:"user_id"`
		UserType string `json:"user_type"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := service.UserServiceInterfaceInstance.UpdateUserType(requestBody.UserID, requestBody.UserType); err != nil {
		log.Printf("UpdateUserType error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User type updated successfully"})
}

func GetAllUsers(c *gin.Context) {
	users, err := service.UserServiceInterfaceInstance.GetAllUsers()
	if err != nil {
		log.Printf("GetAllUsers error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func UpdateUserPhoto(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	currentUser := user.(dao.User)

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxPhotoBytes)

	file, _, err := c.Request.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "photo field required"})
		return
	}
	defer file.Close()

	imageData, err := io.ReadAll(io.LimitReader(file, maxPhotoBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
		return
	}

	photoBase64 := base64.StdEncoding.EncodeToString(imageData)
	if err := clients.UpdateUserPhoto(currentUser.ID, photoBase64); err != nil {
		log.Printf("UpdateUserPhoto error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

func UpdateUserMe(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	currentUser := user.(dao.User)

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := service.UserServiceInterfaceInstance.UpdateUser(currentUser.ID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
