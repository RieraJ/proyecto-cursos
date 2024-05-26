package initializers

import "backend/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Course{}, &models.Teacher{}, &models.CourseTeacher{}, &models.CourseInscription{})
}
