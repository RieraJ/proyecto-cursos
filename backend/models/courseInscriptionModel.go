package models

import "gorm.io/gorm"

type CourseInscription struct {
	gorm.Model
	UserID          uint
	CourseID        uint
	InscriptionDate string `gorm:"type:date"`
	User            User   `gorm:"foreignKey:UserID"`
	Course          Course `gorm:"foreignKey:CourseID"`
}
