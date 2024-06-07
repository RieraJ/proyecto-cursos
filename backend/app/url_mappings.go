package app

import (
	courses "backend/controllers/courses"
	"backend/controllers/users"
	"backend/middleware"
)

func mapUrls() {

	// Rutas para usuarios
	router.POST("/signup", users.Signup)
	router.POST("/login", users.Login)
	//router.GET("/validate", middleware.RequireAuth, users.Validate)
	//router.GET("/users", middleware.RequireAuth, users.GetAllUsers)
	//router.GET("/users/:id", middleware.RequireAuth, users.GetUserByID)
	//router.PUT("/users/:id", middleware.RequireAuth, users.UpdateUserByID)
	//router.DELETE("/users/:id", middleware.RequireAuth, users.DeleteUserByID)

	// Rutas para cursos
	router.POST("/courses", middleware.RequireAuth, middleware.RequireAdmin, courses.CreateCourse)
	router.PUT("/courses/:id", middleware.RequireAuth, middleware.RequireAdmin, courses.UpdateCourseByID)
	router.DELETE("/courses/:id", middleware.RequireAuth, middleware.RequireAdmin, courses.DeleteCourseByID)
	//router.GET("/courses", middleware.RequireAuth, courses.GetAllCourses)
	//router.GET("/courses/:id", middleware.RequireAuth, courses.GetCourseByID)
	router.GET("/users/:id/courses", middleware.RequireAuth, courses.GetUserCourses)
	//router.GET("/courses/:id/users", middleware.RequireAuth, courses.GetCourseUsers)

	// Rutas para inscripciones
	router.POST("/enroll", middleware.RequireAuth, middleware.RequireAdmin, courses.EnrollUser)

}
