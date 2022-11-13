package intializers

import "github.com/harish1907/jwt-go/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.JwtUser{})
}
