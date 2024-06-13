package services

import (
	"backend/clients"
	"backend/dao"
	"backend/dto"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct{}

type authServiceInterface interface {
	Login(LoginRequest dto.User) (string, error)
}

var (
	AuthServiceInterfaceInstance authServiceInterface
)

func init() {
	AuthServiceInterfaceInstance = &authService{}
}

func (s *authService) Login(User dto.User) (string, error) {
	client := dao.User{
		Email:    User.Email,
		Password: User.Password,
	}

	// Verify if user exists
	userDAO, err := clients.SelectUserByEmail(User.Email)
	if err != nil {
		return client.Email, errors.New("invalid email or password")
	}

	// Compare sent in pass with saved user pass hash
	err = bcrypt.CompareHashAndPassword([]byte(userDAO.Password), []byte(client.Password))
	if err != nil {
		return client.Password, errors.New("invalid email or password")
	}

	// Generate a jwt token
	tokenString, err := generateJWT(userDAO.Email, userDAO.ID)
	if err != nil {
		return " ", errors.New("error generating token")
	}

	return tokenString, nil
}

func generateJWT(email string, userId uint) (string, error) {
	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
