package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name    string   `gorm:"not null"`
	Surname string   `gorm:"not null"`
	Courses []Course `gorm:"many2many:course_teachers" json:"courses"`
}
