package clients

import (
	"backend/dao"
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() error {
	// Connect to database
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	err = DB.AutoMigrate(&dao.User{}, &dao.Course{}, &dao.CourseInscription{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	// Add index with maximum length
	err = DB.Exec("CREATE UNIQUE INDEX idx_email ON users (email(50));").Error
	if err != nil {
		return fmt.Errorf("failed to create index: %w", err)
	}
	return nil
}

func CreateUser(user *dao.User) error {
	result := DB.Create(user)
	return result.Error
}

func SelectUserByEmail(email string) (dao.User, error) {
	var user dao.User
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
	}
	return user, nil
}

func SelectUserbyID(id uint) (dao.User, error) {
	var user dao.User
	result := DB.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("user not found")
	}
	return user, result.Error
}

func GetAllUsers() (users []dao.User, err error) {
	var allusers []dao.User
	result := DB.Find(&allusers)
	if result.Error != nil {
		return allusers, result.Error
	}
	return allusers, nil
}

func UpdateUserType(userID uint, userType string) error {
	result := DB.Model(&dao.User{}).
		Where("id = ?", userID).
		Update("user_type", userType)
	return result.Error
}

func CreateCourse(course dao.Course) error {
	result := DB.Create(&course)
	if result.Error != nil {
		log.Println("Error creating course:", result.Error)
	}
	return result.Error
}

func GetAllCourses() ([]dao.Course, error) {
	var courses []dao.Course
	result := DB.Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func ObtainCourseByName(name string) (*dao.Course, error) {
	var course dao.Course
	result := DB.Where("name = ?", name).
		First(&course)
	if result.Error != nil {
		return nil, result.Error
	}
	return &course, nil
}

func ObtainCourseByID(id uint) (*dao.Course, error) {
	var course dao.Course
	result := DB.First(&course, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &course, nil
}

func GetUserInscription(userID uint, courseID uint) (*dao.CourseInscription, error) {
	var inscription dao.CourseInscription
	result := DB.Where("user_id = ? AND course_id = ?", userID, courseID).
		First(&inscription)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &inscription, result.Error
}

func GetUserCourses(userID uint) ([]dao.Course, error) {
	var courses []dao.Course
	result := DB.Table("courses").
		Joins("JOIN course_inscriptions ON courses.id = course_inscriptions.course_id").
		Where("course_inscriptions.user_id = ?", userID).
		Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func UpdateCourseByID(id uint, course dao.Course) error {
	result := DB.Model(&dao.Course{}).Where("id = ?", id).Updates(course)
	if result.Error != nil {
		log.Println("Error updating course:", result.Error)
	}
	return result.Error
}

func DeleteCourseByID(id uint) error {
	result := DB.Delete(&dao.Course{}, id)
	if result.Error != nil {
		log.Println("Error deleting course:", result.Error)
	}
	return result.Error
}

func EnrollUser(inscription dao.CourseInscription) error {
	result := DB.Create(&inscription)
	if result.Error != nil {
		return errors.New("error enrolling user: " + result.Error.Error())
	}
	return nil
}

func SearchCourses(name string) ([]dao.Course, error) {
	nameDB := name
	var courses []dao.Course
	err := DB.Where("name LIKE ?", "%"+nameDB+"%").Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}
