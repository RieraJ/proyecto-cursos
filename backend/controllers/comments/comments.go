package comments

import (
	"backend/dao"
	"backend/dto"
	"backend/services"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const maxCommentContentLen = 2000
const maxImageBytes = 5 << 20 // 5 MB

func CreateComment(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	currentUser := user.(dao.User)

	courseIDStr := c.PostForm("course_id")
	content := c.PostForm("content")

	if courseIDStr == "" || content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	if len(content) > maxCommentContentLen {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content exceeds maximum length"})
		return
	}

	courseID, err := strconv.ParseUint(courseIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
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

	body := dto.CommentRequest{
		UserID:   currentUser.ID,
		CourseID: uint(courseID),
		Content:  content,
		Image:    imageBase64,
	}

	result, err := services.CommentServiceInterfaceInstance.CreateComment(body)
	if err != nil {
		log.Printf("CreateComment error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result.Message})
}

func DeleteCommentByID(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	currentUser := user.(dao.User)

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	isAdmin := currentUser.UserType == "admin"
	err = services.CommentServiceInterfaceInstance.DeleteCommentByID(uint(id), currentUser.ID, isAdmin)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		if err.Error() == "comment not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
			return
		}
		log.Printf("DeleteCommentByID error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func GetUserComments(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	currentUser := user.(dao.User)

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if uint(id) != currentUser.ID && currentUser.UserType != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	comments, err := services.CommentServiceInterfaceInstance.GetUserComments(uint(id))
	if err != nil {
		log.Printf("GetUserComments error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func GetCourseComments(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	comments, err := services.CommentServiceInterfaceInstance.GetCourseComments(uint(id))
	if err != nil {
		log.Printf("GetCourseComments error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
