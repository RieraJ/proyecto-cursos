package middleware

import (
	"backend/clients"
	"backend/dao"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func init() {
	jwtSecret = []byte(os.Getenv("SECRET"))
}

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil || tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token subject"})
		return
	}

	var user dao.User
	if err := clients.DB.Where("email = ?", sub).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("user", user)
	c.Next()
}
