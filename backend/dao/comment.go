package dao

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CourseID uint   `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	Content  string `gorm:"not null"`
	Image    string `gorm:"type:longblob"`
}
