package main

import (
	inscription "backend/controllers/courseInscription"
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
	r.GET("/users", middleware.RequireAuth, users.GetAllUsers)
	r.GET("/users/:id", middleware.RequireAuth, users.GetUserByID)
	r.PUT("/users/:id", middleware.RequireAuth, users.UpdateUserByID)
	r.DELETE("/users/:id", middleware.RequireAuth, users.DeleteUserByID)

	// Rutas para profesores
	r.POST("/teachers", middleware.RequireAuth, teachers.CreateTeacher)
	r.GET("/teachers", middleware.RequireAuth, teachers.GetAllTeachers)
	r.GET("/teachers/:id", middleware.RequireAuth, teachers.GetTeacherByID)

	// Rutas para cursos
	r.POST("/courses", middleware.RequireAuth, courses.CreateCourse)
	r.GET("/courses", middleware.RequireAuth, courses.GetAllCourses)
	r.GET("/courses/:id", middleware.RequireAuth, courses.GetCourseByID)
	r.GET("/users/:id/courses", middleware.RequireAuth, courses.GetUserCourses)
	r.GET("/courses/:id/users", middleware.RequireAuth, courses.GetCourseUsers)

	// Rutas para inscripciones
	r.POST("/enroll", inscription.EnrollStudent)

	r.Run()
}
