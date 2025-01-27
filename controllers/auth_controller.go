package controllers

import (
	"go-gin/dto"
	"go-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Signup(ctx *gin.Context)
}

type authController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &authController{service: service}
}

func (c *authController) Signup(ctx *gin.Context) {
	var req dto.SignupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.Signup(req.Email, req.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}
