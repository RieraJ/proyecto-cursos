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

	// Rutas para cursos
	router.POST("/courses", middleware.RequireAuth, middleware.RequireAdmin, courses.CreateCourse)
	router.PUT("/courses/:id", middleware.RequireAuth, middleware.RequireAdmin, courses.UpdateCourseByID)
	router.DELETE("/courses/:id", middleware.RequireAuth, middleware.RequireAdmin, courses.DeleteCourseByID)
	router.GET("/users/:id/courses", courses.GetUserCourses)
	router.GET("/search-courses", courses.SearchCourses)
	router.GET("/courses", courses.GetAllCourses)

	// Rutas para inscripciones
	router.POST("/enroll", middleware.RequireAuth, courses.EnrollUser)

}
