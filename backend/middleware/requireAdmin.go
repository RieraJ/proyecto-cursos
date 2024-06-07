package middleware

import (
	"backend/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAdmin(c *gin.Context) {
	// Get the user from the context
	user, _ := c.Get("user")

	// Check if the user is an admin
	if user.(dao.User).UserType != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Don't have permissions to access this resource"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Continue if admin
	c.Next()
}
