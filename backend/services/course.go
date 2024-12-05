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
	GetUserCourses(userID uint) ([]dto.Course, error)
	SearchCourses(name string) ([]dto.Course, error)
	GetAllCourses() ([]dto.Course, error)
	GetUserInfo(id uint) (dto.UserInfo, error)
}

var (
	CourseServiceInterfaceInstance courseServiceInterface
)

func init() {
	CourseServiceInterfaceInstance = &courseService{}
}

func (s *courseService) CreateCourse(course dto.Course) (dto.Course, error) {

	var categories []dao.Category
	for _, categoryDTO := range course.Categories {
		category := dao.Category{Name: categoryDTO.Name}
		if err := clients.FindOrCreateCategory(&category); err != nil {
			return course, err
		}
		categories = append(categories, category)
	}

	newCourse := dao.Course{
		Name:         course.Name,
		Description:  course.Description,
		Price:        course.Price,
		Active:       course.Active,
		Instructor:   course.Instructor,
		Length:       course.Length,
		Requirements: course.Requirements,
		Categories:   categories,
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
	// Obtener el curso por ID
	courseDB, err := clients.ObtainCourseByIDWithCategories(id) // Cargar curso con categorías
	if err != nil {
		return dto.Course{}, errors.New("course not found")
	}

	// Actualizar solo los campos no vacíos
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
	if course.Active {
		courseDB.Active = course.Active
	}

	// Actualizar categorías si se proporcionaron
	if len(course.Categories) > 0 {
		var updatedCategories []dao.Category
		for _, categoryDTO := range course.Categories {
			category := dao.Category{Name: categoryDTO.Name}
			if err := clients.FindOrCreateCategory(&category); err != nil {
				return dto.Course{}, err
			}
			updatedCategories = append(updatedCategories, category)
		}
		courseDB.Categories = updatedCategories
	}

	// Guardar los cambios en la base de datos
	if err := clients.UpdateCourseByID(id, *courseDB); err != nil {
		return dto.Course{}, err
	}

	// Convertir a DTO y devolver
	updatedCategoriesDTO := []dto.Category{}
	for _, category := range courseDB.Categories {
		updatedCategoriesDTO = append(updatedCategoriesDTO, dto.Category{
			ID:   category.ID,
			Name: category.Name,
		})
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
		Categories:   updatedCategoriesDTO,
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

func (s *courseService) GetUserCourses(userID uint) ([]dto.Course, error) {
	courses, err := clients.GetUserCoursesWithCategories(userID)
	if err != nil {
		return nil, err
	}
	if len(courses) == 0 {
		return nil, errors.New("no courses found for the user")
	}

	var coursesDTO []dto.Course
	for _, course := range courses {
		var categoriesDTO []dto.Category
		for _, category := range course.Categories {
			categoriesDTO = append(categoriesDTO, dto.Category{
				ID:   category.ID,
				Name: category.Name,
			})
		}

		coursesDTO = append(coursesDTO, dto.Course{
			ID:           course.ID,
			Name:         course.Name,
			Description:  course.Description,
			Price:        course.Price,
			Active:       course.Active,
			Instructor:   course.Instructor,
			Length:       course.Length,
			Requirements: course.Requirements,
			Categories:   categoriesDTO,
		})
	}

	return coursesDTO, nil
}

func (s *courseService) SearchCourses(query string) ([]dto.Course, error) {
	if query == "" {
		return nil, errors.New("search term is empty")
	}

	// Obtener cursos por nombre o por categoría
	coursesDAO, err := clients.SearchCourses(query)
	if err != nil {
		return nil, err
	}

	// Convertir los resultados al formato DTO
	var coursesDTO []dto.Course
	for _, course := range coursesDAO {
		categoriesDTO := []dto.Category{}
		for _, category := range course.Categories {
			categoriesDTO = append(categoriesDTO, dto.Category{
				ID:   category.ID,
				Name: category.Name,
			})
		}

		coursesDTO = append(coursesDTO, dto.Course{
			ID:           course.ID,
			Name:         course.Name,
			Description:  course.Description,
			Price:        course.Price,
			Active:       course.Active,
			Instructor:   course.Instructor,
			Length:       course.Length,
			Requirements: course.Requirements,
			Categories:   categoriesDTO,
		})
	}

	return coursesDTO, nil
}

/*
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
*/

func (s *courseService) GetAllCourses() ([]dto.Course, error) {
	coursesDAO, err := clients.GetAllCoursesWithCategories() // Nueva función para cargar categorías
	if err != nil {
		return nil, err
	}

	var coursesDTO []dto.Course
	for _, course := range coursesDAO {
		var categoriesDTO []dto.Category
		for _, category := range course.Categories {
			categoriesDTO = append(categoriesDTO, dto.Category{
				ID:   category.ID,
				Name: category.Name,
			})
		}

		coursesDTO = append(coursesDTO, dto.Course{
			ID:           course.ID,
			Name:         course.Name,
			Description:  course.Description,
			Price:        course.Price,
			Active:       course.Active,
			Instructor:   course.Instructor,
			Length:       course.Length,
			Requirements: course.Requirements,
			Categories:   categoriesDTO,
		})
	}

	return coursesDTO, nil
}

func (s *courseService) GetUserInfo(id uint) (dto.UserInfo, error) {
	user, err := clients.SelectUserbyID(id)
	if err != nil {
		return dto.UserInfo{}, errors.New("user not found")
	}

	return dto.UserInfo{
		ID:       user.ID,
		Email:    user.Email,
		UserType: user.UserType,
	}, nil
}
