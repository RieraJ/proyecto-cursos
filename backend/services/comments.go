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
	//GetCommentByID(id uint) (*dto.Comment, error)
	DeleteCommentByID(id uint) error
	GetUserComments(userID uint) ([]dto.Comment, error)
	GetCourseComments(courseID uint) ([]dto.Comment, error)
}

var (
	CommentServiceInterfaceInstance CommentServiceInterface = &commentService{}
)

func (s *commentService) CreateComment(comment dto.CommentRequest) (dto.CommentResponse, error) {
	newComment := dao.Comment{
		UserID:   comment.UserID,
		CourseID: comment.CourseID,
		Content:  comment.Content,
		Image:    comment.Image,
	}

	// Check if no image is provided
	if len(newComment.Image) == 0 {
		newComment.Image = nil
	}

	// Check if the comment already exists
	_, err := clients.GetCommentByID(newComment.ID)
	if err == nil {
		return dto.CommentResponse{}, errors.New("comment already exists")
	}

	// Save the comment in the DB
	if err := clients.CreateComment(newComment); err != nil {
		return dto.CommentResponse{Message: "Error whilite creating comment"}, err
	}

	return dto.CommentResponse{Message: "Comment created successfully"}, nil
}

/*
func (s *commentService) GetCommentByID(id uint) (*dto.Comment, error) {
	comment, err := clients.GetCommentByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.Comment{
		ID:       comment.ID,
		UserID:   comment.UserID,
		CourseID: comment.CourseID,
		Content:  comment.Content,
		Image:    comment.Image,
	}, nil
}
*/

func (s *commentService) DeleteCommentByID(id uint) error {
	_, err := clients.GetCommentByID(id)
	if err != nil {
		return errors.New("comment not found")
	}
	// Delete the Comment in the DB
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

	if len(comments) == 0 {
		return nil, errors.New("no comments found")
	}

	var dtoComments []dto.Comment

	for _, comment := range comments {
		dtoComments = append(dtoComments, dto.Comment{
			ID:       comment.ID,
			UserID:   comment.UserID,
			CourseID: comment.CourseID,
			Content:  comment.Content,
			Image:    comment.Image,
		})
	}

	return dtoComments, nil
}

func (s *commentService) GetCourseComments(courseID uint) ([]dto.Comment, error) {
	comments, err := clients.GetCourseComments(courseID)
	if err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return nil, errors.New("no comments found")
	}

	var dtoComments []dto.Comment
	for _, comment := range comments {
		dtoComments = append(dtoComments, dto.Comment{
			ID:       comment.ID,
			UserID:   comment.UserID,
			CourseID: comment.CourseID,
			Content:  comment.Content,
			Image:    comment.Image,
		})
	}

	return dtoComments, nil
}
