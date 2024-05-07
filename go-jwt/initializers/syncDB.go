package initializers

import "github.com/MaaHiN15/go-practice/go-jwt/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}