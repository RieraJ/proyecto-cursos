package models

type CourseTeacher struct {
	CourseID  uint `gorm:"primaryKey"`
	TeacherID uint `gorm:"primaryKey"`
}
