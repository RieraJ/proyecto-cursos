package dao

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CourseID uint   `gorm:"not null"`      // Relación con el curso
	UserID   uint   `gorm:"not null"`      // Relación con el usuario que hizo el comentario
	Content  string `gorm:"not null"`      // Contenido del comentario
	Image    []byte `gorm:"type:longblob"` // Imagen asociada al comentario (opcional)
}
