package services

import (
	"go-gin/dto"
	"go-gin/models"
	"go-gin/repositories"
)

type ItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId, userId uint) (*models.Item, error)
	Create(item dto.CreateItemRequest, userId uint) (*models.Item, error)
	Update(itemId, userId uint, item dto.UpdateItemRequest) (*models.Item, error)
	Delete(itemId, userId uint) error
}

type itemService struct {
	repository repositories.ItemRepository
}

func NewitemRepository(repository repositories.ItemRepository) ItemService {
	return &itemService{repository: repository}
}

func (s *itemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *itemService) FindById(itemId, userId uint) (*models.Item, error) {
	return s.repository.FindById(itemId, userId)
}

func (s *itemService) Create(item dto.CreateItemRequest, userId uint) (*models.Item, error) {
	return s.repository.Create(models.Item{
		Name:        item.Name,
		Price:       item.Price,
		Description: item.Description,
		SoldOut:     false,
		UserID:      userId,
	})
}

func (s *itemService) Update(itemId, userId uint, item dto.UpdateItemRequest) (*models.Item, error) {
	updateItem, err := s.FindById(itemId, userId)
	if err != nil {
		return nil, err
	}
	if item.Name != nil {
		updateItem.Name = *item.Name
	}
	if item.Price != nil {
		updateItem.Price = *item.Price
	}
	if item.Description != nil {
		updateItem.Description = *item.Description
	}
	if item.SoldOut != nil {
		updateItem.SoldOut = *item.SoldOut
	}
	return s.repository.Update(*updateItem)
}

func (s *itemService) Delete(itemId, userId uint) error {
	return s.repository.Delete(itemId, userId)
}
