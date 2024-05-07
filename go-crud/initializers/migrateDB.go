package initializers

import "github.com/MaaHiN15/go-practice/go-crud/models"

func SyncDB() {
	DB.AutoMigrate(&models.Post{})
}