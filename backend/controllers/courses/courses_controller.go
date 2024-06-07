package controllers

/*
func CreateCourse(c *gin.Context) {
	var body struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Active      bool    `json:"active"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	course := models.Course{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Active:      body.Active,
	}

	if initializers.DB.Create(&course).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating the course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course created succesfully"})
}

func GetAllCourses(c *gin.Context) {
	var courses []models.Course

	// Utiliza Preload para cargar las asociaciones
	result := initializers.DB.Preload("Users").Find(&courses)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	var course models.Course

	// Use preload to get the users
	result := initializers.DB.Preload("Users").First(&course, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

// returns the courses that a user is enrolled in
func GetUserCourses(c *gin.Context) {
	userID := c.Param("id")
	var user models.User

	// Find the user and preload the courses
	if initializers.DB.Preload("Courses").First(&user, userID).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"courses": user.Courses})
}

// returns the users that are enrolled in a course
func GetCourseUsers(c *gin.Context) {
	courseID := c.Param("course_id")
	var course models.Course

	// Find the course and preload the users
	if initializers.DB.Preload("Users").First(&course, courseID).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": course.Users})
}
*/
