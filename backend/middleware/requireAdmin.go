package middleware

import (
	"backend/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAdmin(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	currentUser, ok := user.(dao.User)
	if !ok || currentUser.UserType != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	c.Next()
}
