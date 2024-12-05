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

type userService struct{}

type userServiceInterface interface {
	Signup(SignUp dto.SignUpRequest) error
	Login(LoginRequest dto.User) (string, error)
	UpdateUserType(userID uint, userType string) error
	EnrollUser(EnrollUser dto.InscriptionRequest) error
	GetAllUsers() ([]dto.User, error)
	IsAdmin(userID uint) (bool, error)
}

var (
	UserServiceInterfaceInstance userServiceInterface
)

func init() {
	UserServiceInterfaceInstance = &userService{}
}

// Métodos de signup
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

// Métodos de login
func (s *userService) Login(User dto.User) (string, error) {
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

// Métodos de actualización
func (s *userService) UpdateUserType(userID uint, userType string) error {
	// Validar el nuevo tipo de usuario (opcional, si tienes valores predefinidos como "admin", "student")
	validUserTypes := map[string]bool{"admin": true, "student": true}
	if !validUserTypes[userType] {
		return errors.New("invalid user type")
	}

	// Llamar al cliente para actualizar el userType
	err := clients.UpdateUserType(userID, userType)
	if err != nil {
		return err
	}

	return nil
}

// Métodos de administrador
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

// Métodos de inscripción
func (s *userService) EnrollUser(enrollUser dto.InscriptionRequest) error {
	// Verify if user exists
	_, err := clients.SelectUserbyID(enrollUser.UserID)
	if err != nil {
		return err
	}

	// Verify if course exists
	_, err = clients.ObtainCourseByID(enrollUser.CourseID)
	if err != nil {
		return err
	}

	// Verify if user is already enrolled
	inscription, err := clients.GetUserInscription(enrollUser.UserID, enrollUser.CourseID)
	if err != nil {
		return err
	}
	if inscription != nil {
		return errors.New("user is already enrolled in this course")
	}

	// Enroll user
	err = clients.EnrollUser(dao.CourseInscription{
		UserID:   enrollUser.UserID,
		CourseID: enrollUser.CourseID,
	})
	if err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
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
