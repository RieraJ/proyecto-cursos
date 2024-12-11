package comments

import (
	"backend/dto"
	"backend/services"
	"encoding/base64"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	userIDStr := c.PostForm("user_id")
	courseIDStr := c.PostForm("course_id")
	content := c.PostForm("content")

	if courseIDStr == "" || content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// Leer userID y courseID como uint
	userID, _ := strconv.ParseUint(userIDStr, 10, 64)

	courseID, _ := strconv.ParseUint(courseIDStr, 10, 64)

	// Procesar la imagen si existe
	var imageBase64 string
	file, err := c.FormFile("image")
	if err == nil {
		openedFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
			return
		}
		defer openedFile.Close()

		imageData, err := io.ReadAll(openedFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file"})
			return
		}

		imageBase64 = base64.StdEncoding.EncodeToString(imageData)
	}

	// Crear el comentario usando el servicio
	body := dto.CommentRequest{
		UserID:   uint(userID),
		CourseID: uint(courseID),
		Content:  content,
		Image:    imageBase64,
	}

	result, err := services.CommentServiceInterfaceInstance.CreateComment(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result.Message})
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
