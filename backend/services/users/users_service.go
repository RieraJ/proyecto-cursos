package users

import "backend/domain/users"

func Login(request users.LoginRequest) users.LoginResponse {
	// Call the service
	return users.LoginResponse{
		Token: "abcedf123456",
	}
}
