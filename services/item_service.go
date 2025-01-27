package services

import (
	"go-gin/dto"
	"go-gin/models"
	"go-gin/repositories"
)

type ItemService interface {
	FindAll() (*[]models.Item, error)
	FindByID(id uint) (*models.Item, error)
	Create(item dto.CreateItemRequest) (*models.Item, error)
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

func (s *itemService) Create(item dto.CreateItemRequest) (*models.Item, error) {
	return s.repository.Create(models.Item{
		Name:        item.Name,
		Price:       item.Price,
		Description: item.Description,
		SoldOut:     false,
	})
}
