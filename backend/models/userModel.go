package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name               string `gorm:"not null"`
	Surname            string `gorm:"not null"`
	Email              string `gorm:"unique;not null"`
	Password           string `gorm:"not null"`
	UserType           string `gorm:"not null"` // Could be "admin" or "user"
	CourseInscriptions []CourseInscription
}
