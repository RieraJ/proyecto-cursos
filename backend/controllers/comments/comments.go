package comments

import (
	"backend/dto"
	"backend/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var body dto.CommentRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CommentServiceInterfaceInstance.CreateComment(body)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": result.Message})

}

/*
	func GetCommentByID(c *gin.Context) {
		id := c.Param("id")

		idUint, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid comment ID"})
			return
		}

		comment, err := services.CommentServiceInterfaceInstance.GetCommentByID(uint(idUint))
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, comment)
	}
*/
func DeleteCommentByID(c *gin.Context) {
	id := c.Param("id")

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment ID"})
		return
	}

	err = services.CommentServiceInterfaceInstance.DeleteCommentByID(uint(idUint))
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Comment deleted successfully"})
}

func GetUserComments(c *gin.Context) {
	id := c.Param("id")

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	comments, err := services.CommentServiceInterfaceInstance.GetUserComments(uint(idUint))
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, comments)
}

func GetCourseComments(c *gin.Context) {

	id := c.Param("id")

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid course ID"})
		return
	}

	comments, err := services.CommentServiceInterfaceInstance.GetCourseComments(uint(idUint))
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, comments)
}
