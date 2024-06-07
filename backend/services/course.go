package services

import (
	"backend/clients"
	"backend/dao"
	"backend/dto"
	"errors"
)

type courseService struct{}

type courseServiceInterface interface {
	CreateCourse(Course dto.Course) (dto.Course, error)
	UpdateCourseByID(id uint, Course dto.Course) (dto.Course, error)
	DeleteCourseByID(id uint) error
	GetUserCourses(userID uint) ([]dao.Course, error)
}

var (
	CourseServiceInterfaceInstance courseServiceInterface
)

func init() {
	CourseServiceInterfaceInstance = &courseService{}
}

func (s *courseService) CreateCourse(course dto.Course) (dto.Course, error) {
	// Convertir el DTO a un objeto DAO
	newCourse := dao.Course{
		Name:        course.Name,
		Description: course.Description,
		Price:       course.Price,
		Active:      course.Active,
	}

	// Verificar si el curso ya existe
	_, err := clients.ObtainCourseByName(newCourse.Name)
	if err == nil {
		return course, errors.New("course already exists")
	}

	// Guardar el curso en la base de datos
	if err := clients.CreateCourse(newCourse); err != nil {
		return course, err
	}

	return course, nil
}

func (s *courseService) UpdateCourseByID(id uint, course dto.Course) (dto.Course, error) {
	// Obtain the course by ID
	courseDB, err := clients.ObtainCourseByID(id)
	if err != nil {
		return course, errors.New("course not found")
	}

	// Update the course
	courseDB.ID = id
	courseDB.Name = course.Name
	courseDB.Description = course.Description
	courseDB.Price = course.Price
	courseDB.Active = course.Active

	// Save the Course in the DB
	if err := clients.UpdateCourseByID(id, *courseDB); err != nil {
		return course, err
	}

	return dto.Course{
		ID:          courseDB.ID,
		Name:        courseDB.Name,
		Description: courseDB.Description,
		Price:       courseDB.Price,
		Active:      courseDB.Active,
	}, nil
}

func (s *courseService) DeleteCourseByID(id uint) error {
	// Obtain the course by ID
	_, err := clients.ObtainCourseByID(id)
	if err != nil {
		return errors.New("course not found")
	}

	// Delete the Course in the DB
	if err := clients.DeleteCourseByID(id); err !=
		nil {
		return err
	}

	return nil
}

func (s *courseService) GetUserCourses(userID uint) ([]dao.Course, error) {
	courses, err := clients.GetUserCourses(userID)
	if err != nil {
		return nil, err
	}
	if len(courses) == 0 {
		return nil, errors.New("no courses found for the user")
	}

	return courses, nil
}
