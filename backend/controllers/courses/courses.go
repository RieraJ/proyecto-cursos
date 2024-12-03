package controllers

import (
	"backend/clients"
	"backend/dao"
	"backend/dto"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {
	// Get the body of the POST request
	var body dto.Course
	// Unmarshal the JSON body into a new Course struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the CreateCourse service
	result, err := services.CourseServiceInterfaceInstance.CreateCourse(body)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	course, err := clients.ObtainCourseByName(result.Name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	courseID := course.ID
	result.ID = courseID

	// Return the created course in the response
	c.JSON(http.StatusOK, gin.H{"message": "Course successfully created", "course": result})
}

func UpdateCourseByID(c *gin.Context) {
	// Get the body of the PUT request
	var body dto.Course
	// Unmarshal the JSON body into a new Course struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the ID from the URL
	id := c.Param("id")

	// Convert the ID to an unsigned integer
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	// Call the UpdateCourseByID service
	result, err := services.CourseServiceInterfaceInstance.UpdateCourseByID(uint(idUint), body)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return the updated course in the response
	c.JSON(http.StatusOK, gin.H{"message": "Course successfully updated", "course": result})
}

func DeleteCourseByID(c *gin.Context) {
	// Get the ID from the URL
	id := c.Param("id")

	// Convert the ID to an unsigned integer
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	// Call the DeleteCourseByID service
	err = services.CourseServiceInterfaceInstance.DeleteCourseByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return the deleted course in the response
	c.JSON(http.StatusOK, gin.H{"message": "Course successfully deleted"})
}

func GetUserCourses(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := user.(dao.User).ID
	courses, err := services.CourseServiceInterfaceInstance.GetUserCourses(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No courses found for the user"})
		return
	}

	// Convert dao.Course to dto.Course and include categories
	var coursesDTO []dto.Course
	for _, course := range courses {
		var categoriesDTO []dto.Category
		for _, category := range course.Categories {
			categoriesDTO = append(categoriesDTO, dto.Category{
				ID:   category.ID,
				Name: category.Name,
			})
		}

		coursesDTO = append(coursesDTO, dto.Course{
			ID:           course.ID,
			Name:         course.Name,
			Description:  course.Description,
			Price:        course.Price,
			Active:       course.Active,
			Instructor:   course.Instructor,
			Length:       course.Length,
			Requirements: course.Requirements,
			Categories:   categoriesDTO,
		})
	}

	c.JSON(http.StatusOK, gin.H{"courses": coursesDTO})
}

func SearchCourses(c *gin.Context) {
	name := c.Query("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "search term is empty"})
		return
	}

	courses, err := services.CourseServiceInterfaceInstance.SearchCourses(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(courses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no courses found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"courses": courses})
}

func GetAllCourses(c *gin.Context) {
	courses, err := services.CourseServiceInterfaceInstance.GetAllCourses()
	if err != nil {
		if err.Error() == "no courses found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Formatear la respuesta para incluir categorías
	var formattedCourses []dto.Course
	for _, course := range courses {
		var categoriesDTO []dto.Category
		for _, category := range course.Categories {
			categoriesDTO = append(categoriesDTO, dto.Category{
				ID:   category.ID,
				Name: category.Name,
			})
		}

		formattedCourses = append(formattedCourses, dto.Course{
			ID:           course.ID,
			Name:         course.Name,
			Description:  course.Description,
			Price:        course.Price,
			Active:       course.Active,
			Instructor:   course.Instructor,
			Length:       course.Length,
			Requirements: course.Requirements,
			Categories:   categoriesDTO,
		})
	}

	c.JSON(http.StatusOK, gin.H{"courses": formattedCourses})
}

func GetUserInfo(c *gin.Context) {
	// Obtener el ID del usuario desde el contexto (donde se estableció en RequireAuth)
	user, _ := c.Get("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Obtener el ID del usuario
	userID := user.(dao.User).ID

	// Llamar al servicio para obtener la información del usuario
	userInfo, err := services.CourseServiceInterfaceInstance.GetUserInfo(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver la información del usuario en la respuesta
	c.JSON(http.StatusOK, gin.H{"userInfo": userInfo})
}
