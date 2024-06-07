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

func (s *inscriptionService) EnrollUser(EnrollUser dto.InscriptionRequest) error {
	// Verify if user exists
	_, err := clients.SelectUserbyID(EnrollUser.UserID)
	if err != nil {
		return err
	}

	// Verify if course exists
	_, err = clients.ObtainCourseByID(EnrollUser.CourseID)
	if err != nil {
		return err
	}

	// Verify if user is already enrolled
	inscription, err := clients.GetUserInscription(EnrollUser.UserID, EnrollUser.CourseID)
	if err != nil {
		return err
	}
	if inscription != nil {
		return errors.New("user is already enrolled in this course")
	}

	// Enroll user
	err = clients.EnrollUser(dao.CourseInscription{
		UserID:   EnrollUser.UserID,
		CourseID: EnrollUser.CourseID,
	})
	if err != nil {
		return err
	}

	return nil

}
