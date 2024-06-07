package services

import (
	"backend/clients"
	"backend/dto"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct{}

type authServiceInterface interface {
	Login(LoginRequest dto.LoginRequest) (dto.LoginResponse, error)
}

var (
	AuthServiceInterfaceInstance authServiceInterface
)

func init() {
	AuthServiceInterfaceInstance = &authService{}
}

func (s *authService) Login(LoginRequest dto.LoginRequest) (dto.LoginResponse, error) {
	var response dto.LoginResponse

	// Verify if user exists
	userDAO, err := clients.SelectUserByEmail(LoginRequest.Email)
	if err != nil {
		return response, errors.New("invalid email or password")
	}

	// Compare sent in pass with saved user pass hash
	err = bcrypt.CompareHashAndPassword([]byte(userDAO.Password), []byte(LoginRequest.Password))
	if err != nil {
		return response, errors.New("invalid email or password")
	}

	// Generate a jwt token
	tokenString, err := generateJWT(userDAO.Email)
	if err != nil {
		return response, errors.New("error generating token")
	}

	response.Token = tokenString
	return response, nil
}

func generateJWT(email string) (string, error) {
	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
