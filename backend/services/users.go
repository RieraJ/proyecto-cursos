package services

import (
	"backend/clients"
	"backend/dao"
	"backend/dto"
	"errors"
	"net/mail"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

type userServiceInterface interface {
	Signup(SignUp dto.SignUpRequest) error
	Login(LoginRequest dto.User) (string, error)
	UpdateUserType(userID uint, userType string) error
	UpdateUser(userID uint, req dto.UpdateUserRequest) error
	EnrollUser(EnrollUser dto.InscriptionRequest) error
	GetAllUsers() ([]dto.User, error)
	IsAdmin(userID uint) (bool, error)
}

var (
	UserServiceInterfaceInstance userServiceInterface
	jwtSecret                    []byte
)

func init() {
	UserServiceInterfaceInstance = &userService{}
	jwtSecret = []byte(os.Getenv("SECRET"))
}

func (s *userService) Signup(signUp dto.SignUpRequest) error {
	if err := validateSignupInput(signUp); err != nil {
		return err
	}

	_, err := clients.SelectUserByEmail(signUp.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := hashPassword(signUp.Password)
	if err != nil {
		return err
	}

	newUser := &dao.User{
		Email:    signUp.Email,
		Password: hashedPassword,
		Name:     signUp.Name,
		Surname:  signUp.Surname,
		UserType: "student",
	}

	return clients.CreateUser(newUser)
}

func (s *userService) Login(User dto.User) (string, error) {
	userDAO, err := clients.SelectUserByEmail(User.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDAO.Password), []byte(User.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	tokenString, err := generateJWT(userDAO.Email, userDAO.ID)
	if err != nil {
		return "", errors.New("error generating token")
	}

	return tokenString, nil
}

func (s *userService) UpdateUserType(userID uint, userType string) error {
	validUserTypes := map[string]bool{"admin": true, "student": true}
	if !validUserTypes[userType] {
		return errors.New("invalid user type")
	}

	return clients.UpdateUserType(userID, userType)
}

func (s *userService) UpdateUser(userID uint, req dto.UpdateUserRequest) error {
	updates := map[string]interface{}{}

	if req.Name != "" {
		if len(req.Name) > 100 {
			return errors.New("name exceeds maximum length")
		}
		updates["name"] = req.Name
	}
	if req.Surname != "" {
		if len(req.Surname) > 100 {
			return errors.New("surname exceeds maximum length")
		}
		updates["surname"] = req.Surname
	}
	if req.Email != "" {
		if _, err := mail.ParseAddress(req.Email); err != nil {
			return errors.New("invalid email format")
		}
		existing, err := clients.SelectUserByEmail(req.Email)
		if err == nil && existing.ID != userID {
			return errors.New("email already in use")
		}
		updates["email"] = req.Email
	}
	if req.Password != "" {
		if len(req.Password) < 8 {
			return errors.New("password must be at least 8 characters")
		}
		hashed, err := hashPassword(req.Password)
		if err != nil {
			return err
		}
		updates["password"] = hashed
	}

	if len(updates) == 0 {
		return nil
	}

	return clients.UpdateUser(userID, updates)
}

func (s *userService) IsAdmin(userID uint) (bool, error) {
	user, err := clients.SelectUserbyID(userID)
	if err != nil {
		return false, err
	}

	return user.UserType == "admin", nil
}

func (s *userService) GetAllUsers() ([]dto.User, error) {
	users, err := clients.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var usersDTO []dto.User
	for _, user := range users {
		usersDTO = append(usersDTO, dto.User{
			ID:       user.ID,
			Email:    user.Email,
			Name:     user.Name,
			Surname:  user.Surname,
			UserType: user.UserType,
		})
	}

	return usersDTO, nil
}

func (s *userService) EnrollUser(enrollUser dto.InscriptionRequest) error {
	_, err := clients.SelectUserbyID(enrollUser.UserID)
	if err != nil {
		return err
	}

	_, err = clients.ObtainCourseByID(enrollUser.CourseID)
	if err != nil {
		return err
	}

	inscription, err := clients.GetUserInscription(enrollUser.UserID, enrollUser.CourseID)
	if err != nil {
		return err
	}
	if inscription != nil {
		return errors.New("user is already enrolled in this course")
	}

	return clients.EnrollUser(dao.CourseInscription{
		UserID:   enrollUser.UserID,
		CourseID: enrollUser.CourseID,
	})
}

func validateSignupInput(signUp dto.SignUpRequest) error {
	if signUp.Name == "" || len(signUp.Name) > 100 {
		return errors.New("name is required and must be under 100 characters")
	}
	if signUp.Surname == "" || len(signUp.Surname) > 100 {
		return errors.New("surname is required and must be under 100 characters")
	}
	if _, err := mail.ParseAddress(signUp.Email); err != nil {
		return errors.New("invalid email format")
	}
	if len(signUp.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func generateJWT(email string, userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	return token.SignedString(jwtSecret)
}
