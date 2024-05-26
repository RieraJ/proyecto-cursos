package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Price              float64
	Active             bool
	Name               string    `gorm:"not null"`
	Description        string    `gorm:"type:text"`
	Teachers           []Teacher `gorm:"many2many:course_teachers;"`
	CourseInscriptions []CourseInscription
}
