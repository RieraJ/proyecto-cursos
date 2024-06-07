package dao

type CourseInscription struct {
	UserID   uint `gorm:"primaryKey"`
	CourseID uint `gorm:"primaryKey"`
}
