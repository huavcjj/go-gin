package main

import (
	"go-gin/controllers"
	"go-gin/infra"
	"go-gin/repositories"
	"go-gin/services"

	_ "go-gin/migrations"

	"github.com/gin-gonic/gin"
)

func main() {
	db := infra.SetupDB()
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewitemRepository(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	itemRouter := r.Group("/items")
	itemRouter.GET("", itemController.FindAll)
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)

	r.Run("localhost:8080")
}
