package main

import (
	"go-gin/controllers"
	"go-gin/infra"
	"go-gin/middlewares"
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

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")

	itemRouter.GET("", itemController.FindAll)
	itemRouterWithAuth.GET("/:id", itemController.FindById)
	itemRouterWithAuth.POST("", itemController.Create)
	itemRouterWithAuth.PUT("/:id", itemController.Update)
	itemRouterWithAuth.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)
	r.Run("localhost:8080")
}
