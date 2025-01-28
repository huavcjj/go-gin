package controllers

import (
	"go-gin/dto"
	"go-gin/models"
	"go-gin/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type itemController struct {
	service services.ItemService
}

func NewItemController(service services.ItemService) ItemController {
	return &itemController{service: service}
}

func (c *itemController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, items)
}

func (c *itemController) FindById(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userId := user.(*models.User).ID

	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := c.service.FindById(uint(itemId), userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}

func (c *itemController) Create(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userId := user.(*models.User).ID

	var req dto.CreateItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newItem, err := c.service.Create(req, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": newItem})
}

func (c *itemController) Update(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userId := user.(*models.User).ID

	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req dto.UpdateItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedItem, err := c.service.Update(uint(itemId), userId, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": updatedItem})
}

func (c *itemController) Delete(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userId := user.(*models.User).ID

	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.service.Delete(uint(itemId), userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}
