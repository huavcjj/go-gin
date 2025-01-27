package migrations

import (
	"go-gin/infra"
	"go-gin/models"
)

func init() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}, &models.User{}); err != nil {
		panic("failed to migrate")
	}
}
