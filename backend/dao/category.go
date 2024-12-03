package dao

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name    string   `gorm:"unique"`                      // Category name
	Courses []Course `gorm:"many2many:course_categories"` // Courses in the category
}
