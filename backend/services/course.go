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
	SearchCourses(name string) ([]dto.Course, error)
	GetAllCourses() ([]dao.Course, error)
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
		Name:         course.Name,
		Description:  course.Description,
		Price:        course.Price,
		Active:       course.Active,
		Instructor:   course.Instructor,
		Length:       course.Length,
		Requirements: course.Requirements,
		Image:        course.Image,
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

	// Update only non-empty fields
	if course.Name != "" {
		courseDB.Name = course.Name
	}
	if course.Description != "" {
		courseDB.Description = course.Description
	}
	if course.Price != 0 {
		courseDB.Price = course.Price
	}
	if course.Instructor != "" {
		courseDB.Instructor = course.Instructor
	}
	if course.Length != "" {
		courseDB.Length = course.Length
	}
	if course.Requirements != "" {
		courseDB.Requirements = course.Requirements
	}
	if course.Image != "" {
		courseDB.Image = course.Image
	}
	if course.Active {
		courseDB.Active = course.Active
	}

	// Save the Course in the DB
	if err := clients.UpdateCourseByID(id, *courseDB); err != nil {
		return course, err
	}

	return dto.Course{
		ID:           courseDB.ID,
		Name:         courseDB.Name,
		Description:  courseDB.Description,
		Price:        courseDB.Price,
		Active:       courseDB.Active,
		Instructor:   courseDB.Instructor,
		Length:       courseDB.Length,
		Requirements: courseDB.Requirements,
		Image:        courseDB.Image,
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

func (s *courseService) SearchCourses(name string) ([]dto.Course, error) {
	if name == "" {
		return nil, errors.New("search term is empty")
	}

	coursesDAO, err := clients.SearchCourses(name)
	if err != nil {
		return nil, err
	}

	var coursesDTO []dto.Course
	for _, course := range coursesDAO {
		coursesDTO = append(coursesDTO, dto.Course{
			ID:           course.ID,
			Name:         course.Name,
			Description:  course.Description,
			Price:        course.Price,
			Active:       course.Active,
			Instructor:   course.Instructor,
			Length:       course.Length,
			Requirements: course.Requirements,
			Image:        course.Image,
		})
	}

	return coursesDTO, nil
}

func (s *courseService) GetAllCourses() ([]dao.Course, error) {
	courses, err := clients.GetAllCourses()
	if err != nil {
		return nil, err
	}
	if len(courses) == 0 {
		return nil, errors.New("no courses found")
	}

	return courses, nil
}
