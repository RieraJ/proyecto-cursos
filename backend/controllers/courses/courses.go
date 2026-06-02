package controllers

import (
	"backend/clients"
	"backend/dao"
	"backend/dto"
	"backend/services"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const maxImageBytes = 5 << 20 // 5 MB

func CreateCourse(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")
	instructor := c.PostForm("instructor")
	length := c.PostForm("length")
	requirements := c.PostForm("requirements")
	categoriesJSON := c.PostForm("categories")

	if name == "" || description == "" || priceStr == "" || instructor == "" || length == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price format"})
		return
	}

	var categories []dto.Category
	if categoriesJSON != "" {
		if err := json.Unmarshal([]byte(categoriesJSON), &categories); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid categories format"})
			return
		}
	}

	var imageBase64 string
	file, err := c.FormFile("image")
	if err == nil {
		if file.Size > maxImageBytes {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "Image exceeds 5 MB limit"})
			return
		}
		openedFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing image"})
			return
		}
		defer openedFile.Close()

		imageData, err := io.ReadAll(io.LimitReader(openedFile, maxImageBytes))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing image"})
			return
		}
		imageBase64 = base64.StdEncoding.EncodeToString(imageData)
	}

	body := dto.Course{
		Name:         name,
		Description:  description,
		Price:        price,
		Active:       true,
		Instructor:   instructor,
		Length:       length,
		Requirements: requirements,
		Categories:   categories,
		Image:        imageBase64,
	}

	result, err := services.CourseServiceInterfaceInstance.CreateCourse(body)
	if err != nil {
		log.Printf("CreateCourse error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	course, err := clients.ObtainCourseByName(result.Name)
	if err != nil {
		log.Printf("ObtainCourseByName error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve created course"})
		return
	}

	result.ID = course.ID
	c.JSON(http.StatusOK, gin.H{"message": "Course successfully created", "course": result})
}

func UpdateCourseByID(c *gin.Context) {
	var body dto.Course
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	result, err := services.CourseServiceInterfaceInstance.UpdateCourseByID(uint(idUint), body)
	if err != nil {
		log.Printf("UpdateCourseByID error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course successfully updated", "course": result})
}

func DeleteCourseByID(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	if err := services.CourseServiceInterfaceInstance.DeleteCourseByID(uint(idUint)); err != nil {
		log.Printf("DeleteCourseByID error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course successfully deleted"})
}

func GetUserCourses(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID := user.(dao.User).ID
	courses, err := services.CourseServiceInterfaceInstance.GetUserCourses(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No courses found for the user"})
		return
	}

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
			Image:        course.Image,
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
		log.Printf("SearchCourses error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search courses"})
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
			c.JSON(http.StatusNotFound, gin.H{"error": "No courses found"})
		} else {
			log.Printf("GetAllCourses error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve courses"})
		}
		return
	}

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
			Image:        course.Image,
		})
	}

	c.JSON(http.StatusOK, gin.H{"courses": formattedCourses})
}

func GetUserInfo(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID := user.(dao.User).ID

	userInfo, err := services.CourseServiceInterfaceInstance.GetUserInfo(userID)
	if err != nil {
		log.Printf("GetUserInfo error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userInfo": userInfo})
}
