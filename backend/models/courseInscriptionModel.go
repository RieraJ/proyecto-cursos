package models

import "time"

type CourseInscription struct {
	UserID          uint      `gorm:"not null"`
	CourseID        uint      `gorm:"not null"`
	InscriptionDate time.Time `gorm:"not null"`
}
