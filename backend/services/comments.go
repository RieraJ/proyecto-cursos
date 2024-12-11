package services

import (
	"backend/clients"
	"backend/dao"
	"backend/dto"
	"errors"
)

type commentService struct{}

type CommentServiceInterface interface {
	CreateComment(comment dto.CommentRequest) (dto.CommentResponse, error)
	DeleteCommentByID(id uint) error
	GetUserComments(userID uint) ([]dto.Comment, error)
	GetCourseComments(courseID uint) ([]dto.Comment, error)
}

var (
	CommentServiceInterfaceInstance CommentServiceInterface = &commentService{}
)

func (s *commentService) CreateComment(comment dto.CommentRequest) (dto.CommentResponse, error) {
	// Obtain the user ID
	_, err := clients.SelectUserbyID(comment.UserID)
	if err != nil {
		return dto.CommentResponse{Message: "User not found"}, err
	}

	// Obtain the course inscription
	_, err = clients.GetCourseInscriptionByUserIDAndCourseID(comment.UserID, comment.CourseID)

	if err != nil {
		return dto.CommentResponse{Message: "User is not enrolled in the course"}, err
	}

	newComment := dao.Comment{
		UserID:   comment.UserID,
		CourseID: comment.CourseID,
		Content:  comment.Content,
		Image:    comment.Image,
	}

	// Save the comment in the DB
	if err := clients.CreateComment(newComment); err != nil {
		return dto.CommentResponse{Message: "Error while creating comment"}, err
	}

	return dto.CommentResponse{Message: "Comment created successfully"}, nil
}

func (s *commentService) DeleteCommentByID(id uint) error {
	_, err := clients.GetCommentByID(id)
	if err != nil {
		return errors.New("comment not found")
	}

	// Delete the comment in the DB
	if err := clients.DeleteCommentByID(id); err != nil {
		return err
	}

	return nil
}

func (s *commentService) GetUserComments(userID uint) ([]dto.Comment, error) {
	comments, err := clients.GetUserComments(userID)
	if err != nil {
		return nil, err
	}

	return transformCommentsToDTO(comments), nil
}

func (s *commentService) GetCourseComments(courseID uint) ([]dto.Comment, error) {
	comments, err := clients.GetCourseComments(courseID)
	if err != nil {
		return nil, err
	}

	return transformCommentsToDTO(comments), nil
}

// Helper function to transform comments to DTO
func transformCommentsToDTO(comments []dao.Comment) []dto.Comment {
	var dtoComments []dto.Comment

	for _, comment := range comments {
		var imageBase64 string
		if len(comment.Image) > 0 {
			imageBase64 = comment.Image
		}

		dtoComments = append(dtoComments, dto.Comment{
			ID:       comment.ID,
			UserID:   comment.UserID,
			CourseID: comment.CourseID,
			Content:  comment.Content,
			Image:    imageBase64,
		})
	}

	return dtoComments
}
