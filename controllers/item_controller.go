package controllers

import (
	"go-gin/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
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
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := c.service.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}
