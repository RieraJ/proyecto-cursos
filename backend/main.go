package main

import (
	"backend/initializers"
	"backend/middleware"

	"backend/controllers/users"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", users.Signup)
	r.POST("/login", users.Login)
	r.GET("/validate", middleware.RequireAuth, users.Validate)

	r.Run()

}
