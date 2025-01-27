package main

import (
	"go-gin/controllers"
	"go-gin/infra"
	"go-gin/repositories"
	"go-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := infra.SetupDB()
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewitemRepository(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.POST("/items", itemController.Create)
	r.PUT("/items/:id", itemController.Update)
	r.DELETE("/items/:id", itemController.Delete)

	r.Run("localhost:8080")
}
