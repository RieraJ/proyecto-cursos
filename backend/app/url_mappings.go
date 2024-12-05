package app

import (
	comments "backend/controllers/comments"
	courses "backend/controllers/courses"
	"backend/controllers/users"
	"backend/middleware"
)

func mapUrls() {

	// Rutas para usuarios
	router.POST("/signup", users.Signup)
	router.POST("/login", users.Login)
	router.GET("/user-info", middleware.RequireAuth, courses.GetUserInfo)
	router.GET("/users", middleware.RequireAuth, middleware.RequireAdmin, users.GetAllUsers)
	router.PUT("/update-user-type", middleware.RequireAuth, middleware.RequireAdmin, users.UpdateUserType)

	// Rutas para cursos
	router.POST("/courses", middleware.RequireAuth, middleware.RequireAdmin, courses.CreateCourse)
	router.PUT("/courses/:id", middleware.RequireAuth, middleware.RequireAdmin, courses.UpdateCourseByID)
	router.DELETE("/courses/:id", middleware.RequireAuth, middleware.RequireAdmin, courses.DeleteCourseByID)
	router.GET("/users/:id/courses", middleware.RequireAuth, courses.GetUserCourses)
	router.GET("/search-courses", courses.SearchCourses)
	router.GET("/courses", courses.GetAllCourses)

	// Rutas para inscripciones
	router.POST("/enroll", middleware.RequireAuth, courses.EnrollUser)

	// Rutas para comentarios
	router.POST("/comments", middleware.RequireAuth, comments.CreateComment)
	router.DELETE("/comments/:id", middleware.RequireAuth, comments.DeleteCommentByID)
	router.GET("/users/:id/comments", middleware.RequireAuth, comments.GetUserComments)
	router.GET("/courses/:id/comments", comments.GetCourseComments)
}
