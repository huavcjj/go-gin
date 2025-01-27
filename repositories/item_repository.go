package repositories

import (
	"errors"
	"go-gin/models"
)

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(id uint) (*models.Item, error)
	Create(item models.Item) (*models.Item, error)
}

type itemRepository struct {
	items []models.Item
}

func NewItemRepository(items []models.Item) ItemRepository {
	return &itemRepository{items: items}
}

func (r *itemRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *itemRepository) FindById(id uint) (*models.Item, error) {
	for _, item := range r.items {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *itemRepository) Create(item models.Item) (*models.Item, error) {
	item.ID = uint(len(r.items) + 1)
	r.items = append(r.items, item)
	return &item, nil
}
