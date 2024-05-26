package main

import (
	courses "backend/controllers/courses"
	teachers "backend/controllers/teachers"
	"backend/controllers/users"
	"backend/initializers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	// Rutas para usuarios
	r.POST("/signup", users.Signup)
	r.POST("/login", users.Login)
	r.GET("/validate", middleware.RequireAuth, users.Validate)

	// Rutas para profesores
	r.POST("/teachers", middleware.RequireAuth, teachers.CreateTeacher)

	// Rutas para cursos
	r.POST("/courses", middleware.RequireAuth, courses.CreateCourse)

	r.Run()
}
