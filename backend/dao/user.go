package dao

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Surname  string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;idx_email;unique"`
	Password string `gorm:"not null"`
	UserType string `gorm:"not null"`
}
