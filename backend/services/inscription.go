package services

import (
	"backend/clients"
	"backend/dao"
	"backend/dto"
	"errors"
)

type inscriptionService struct{}

type inscriptionServiceInterface interface {
	EnrollUser(EnrollUser dto.InscriptionRequest) error
}

var (
	InscriptionServiceInterfaceInstance inscriptionServiceInterface
)

func init() {
	InscriptionServiceInterfaceInstance = &inscriptionService{}
}

func (s *inscriptionService) EnrollUser(enrollUser dto.InscriptionRequest) error {
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
