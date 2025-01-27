package services

import (
	"go-gin/models"
	"go-gin/repositories"
)

type ItemService interface {
	FindAll() (*[]models.Item, error)
	FindByID(id uint) (*models.Item, error)
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

func (s *itemService) FindByID(id uint) (*models.Item, error) {
	return s.repository.FindById(id)
}
