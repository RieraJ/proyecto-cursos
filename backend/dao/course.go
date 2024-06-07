package dao

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Price       float64 `gorm:"not null"`
	Active      bool    `gorm:"not null"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
}
