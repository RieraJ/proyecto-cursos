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
	DeleteCommentByID(id uint, callerID uint, isAdmin bool) error
	GetUserComments(userID uint) ([]dto.Comment, error)
	GetCourseComments(courseID uint) ([]dto.Comment, error)
}

var (
	CommentServiceInterfaceInstance CommentServiceInterface = &commentService{}
)

func (s *commentService) CreateComment(comment dto.CommentRequest) (dto.CommentResponse, error) {
	_, err := clients.SelectUserbyID(comment.UserID)
	if err != nil {
		return dto.CommentResponse{Message: "User not found"}, err
	}

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

	if err := clients.CreateComment(newComment); err != nil {
		return dto.CommentResponse{Message: "Error while creating comment"}, err
	}

	return dto.CommentResponse{Message: "Comment created successfully"}, nil
}

func (s *commentService) DeleteCommentByID(id uint, callerID uint, isAdmin bool) error {
	comment, err := clients.GetCommentByID(id)
	if err != nil {
		return errors.New("comment not found")
	}

	if !isAdmin && comment.UserID != callerID {
		return errors.New("forbidden")
	}

	return clients.DeleteCommentByID(id)
}

func (s *commentService) GetUserComments(userID uint) ([]dto.Comment, error) {
	comments, err := clients.GetUserComments(userID)
	if err != nil {
		return nil, err
	}

	return enrichCommentsWithUserInfo(comments)
}

func (s *commentService) GetCourseComments(courseID uint) ([]dto.Comment, error) {
	comments, err := clients.GetCourseComments(courseID)
	if err != nil {
		return nil, err
	}

	return enrichCommentsWithUserInfo(comments)
}

func enrichCommentsWithUserInfo(comments []dao.Comment) ([]dto.Comment, error) {
	userIDSet := make(map[uint]bool)
	for _, c := range comments {
		userIDSet[c.UserID] = true
	}
	userIDs := make([]uint, 0, len(userIDSet))
	for id := range userIDSet {
		userIDs = append(userIDs, id)
	}

	users, err := clients.GetUsersByIDs(userIDs)
	if err != nil {
		return nil, err
	}

	userMap := make(map[uint]dao.User)
	for _, u := range users {
		userMap[u.ID] = u
	}

	var dtoComments []dto.Comment
	for _, comment := range comments {
		user := userMap[comment.UserID]
		dtoComments = append(dtoComments, dto.Comment{
			ID:          comment.ID,
			UserID:      comment.UserID,
			CourseID:    comment.CourseID,
			Content:     comment.Content,
			Image:       comment.Image,
			UserName:    user.Name,
			UserSurname: user.Surname,
			UserImage:   user.Image,
		})
	}

	return dtoComments, nil
}
