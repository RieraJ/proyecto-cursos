package dto

type Comment struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	CourseID uint   `json:"course_id"`
	Content  string `json:"content"`
	Image    []byte `json:"image"`
}

type CommentRequest struct {
	UserID   uint   `json:"user_id"`
	CourseID uint   `json:"course_id"`
	Content  string `json:"content"`
	Image    []byte `json:"image"`
}

type CommentResponse struct {
	Message string `json:"message"`
}
