package dto

type Comment struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	CourseID uint   `json:"course_id"`
	Content  string `json:"content"`
	Image    string `json:"image"`
}

type CommentRequest struct {
	UserID   uint   `json:"user_id" form:"user_id"`
	CourseID uint   `json:"course_id" form:"course_id"`
	Content  string `json:"content" form:"content"`
	Image    string `json:"image" form:"image"`
}

type CommentResponse struct {
	Message string `json:"message"`
}
