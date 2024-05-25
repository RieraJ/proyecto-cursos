package initializers

import "backend/models"

func SyncDatabase() {
	// Migrate the schema
	DB.AutoMigrate(&models.User{})

}
