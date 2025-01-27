package main

import (
	"go-gin/controllers"
	"go-gin/models"
	"go-gin/repositories"
	"go-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: false},
		{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: true},
	}

	itemRepository := repositories.NewItemRepository(items)
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

	r.Run("localhost:8080")
}
