package dto

type UserInfo struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	UserType string `json:"userType"`
}
