package comments

import (
	"backend/dto"
	"backend/services"
	"encoding/base64"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var body dto.CommentRequest

	// Parse form data
	userID, _ := strconv.ParseUint(c.PostForm("user_id"), 10, 64)
	courseID, _ := strconv.ParseUint(c.PostForm("course_id"), 10, 64)
	content := c.PostForm("content")

	// Process image file
	file, err := c.FormFile("image")
	var imageData []byte
	if err == nil {
		openedFile, err := file.Open()
		if err != nil {
			c.JSON(400, gin.H{"error": "Error opening file"})
			return
		}
		defer openedFile.Close()

		imageData, err = io.ReadAll(openedFile)
		if err != nil {
			c.JSON(400, gin.H{"error": "Error reading file"})
			return
		}
	}
	imageBase64 := base64.StdEncoding.EncodeToString(imageData)
	body = dto.CommentRequest{
		UserID:   uint(userID),
		CourseID: uint(courseID),
		Content:  content,
		Image:    imageBase64,
	}

	result, err := services.CommentServiceInterfaceInstance.CreateComment(body)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": result.Message})
}

func DeleteCommentByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment ID"})
		return
	}

	err = services.CommentServiceInterfaceInstance.DeleteCommentByID(uint(id))
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Comment deleted successfully"})
}

func GetUserComments(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	comments, err := services.CommentServiceInterfaceInstance.GetUserComments(uint(id))
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, comments)
}

func GetCourseComments(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid course ID"})
		return
	}

	comments, err := services.CommentServiceInterfaceInstance.GetCourseComments(uint(id))
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, comments)
}
