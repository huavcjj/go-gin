package main

import (
	"go-gin/infra"
	"go-gin/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("failed to migrate")
	}
}
