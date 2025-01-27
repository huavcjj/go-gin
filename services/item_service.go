package services

import (
	"go-gin/dto"
	"go-gin/models"
	"go-gin/repositories"
)

type ItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(id uint) (*models.Item, error)
	Create(item dto.CreateItemRequest) (*models.Item, error)
	Update(id uint, item dto.UpdateItemRequest) (*models.Item, error)
	Delete(id uint) error
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

func (s *itemService) FindById(id uint) (*models.Item, error) {
	return s.repository.FindById(id)
}

func (s *itemService) Create(item dto.CreateItemRequest) (*models.Item, error) {
	return s.repository.Create(models.Item{
		Name:        item.Name,
		Price:       item.Price,
		Description: item.Description,
		SoldOut:     false,
	})
}

func (s *itemService) Update(id uint, item dto.UpdateItemRequest) (*models.Item, error) {
	updateItem, err := s.FindById(id)
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

func (s *itemService) Delete(id uint) error {
	return s.repository.Delete(id)
}
