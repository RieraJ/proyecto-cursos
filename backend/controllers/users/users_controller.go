package users

/*
func Login(c *gin.Context) {
	// Get the email/pass of request body
	var body domain.LoginRequest

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})

		return
	}
	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid email/password"})
		return
	}

	// Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid email/password"})
		return
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while generating the token"})
		return
	}

	// Respond with the token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 60*60*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{"message": user})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUserByID(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
		Name     string `json:"name,omitempty"`
		Surname  string `json:"surname,omitempty"`
		UserType string `json:"userType,omitempty"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if body.Email != "" {
		user.Email = body.Email
	}
	if body.Name != "" {
		user.Name = body.Name
	}
	if body.Surname != "" {
		user.Surname = body.Surname
	}
	if body.UserType != "" {
		user.UserType = body.UserType
	}

	if body.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing the password"})
			return
		}
		user.Password = string(hash)
	}

	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	initializers.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
*/
