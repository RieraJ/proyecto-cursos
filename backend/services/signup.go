package services

import (
	"backend/clients"
	"backend/dao"
	"backend/dto"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

type userServiceInterface interface {
	Signup(SignUp dto.SignUpRequest) error
	UpdateUserType(userID uint, userType string) error
}

var (
	UserServiceInterfaceInstance userServiceInterface
)

func init() {
	UserServiceInterfaceInstance = &userService{}
}

func (s *userService) Signup(signUp dto.SignUpRequest) error {
	// Verificar si el usuario ya existe
	_, err := clients.SelectUserByEmail(signUp.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash de la contraseña
	hashedPassword, err := hashPassword(signUp.Password)
	if err != nil {
		return err
	}

	// Crear un nuevo usuario
	newUser := &dao.User{
		Email:    signUp.Email,
		Password: hashedPassword,
		Name:     signUp.Name,
		Surname:  signUp.Surname,
		UserType: "student",
	}

	// Guardar el usuario en la base de datos
	if err := clients.CreateUser(newUser); err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateUserType(userID uint, userType string) error {
	return clients.UpdateUserType(userID, userType)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
