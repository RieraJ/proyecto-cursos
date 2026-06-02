package dto

type UserInfo struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	UserType string `json:"userType"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Image    string `json:"image"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
