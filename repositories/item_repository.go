package repositories

import (
	"go-gin/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(id uint) (*models.Item, error)
	Create(item models.Item) (*models.Item, error)
	Update(item models.Item) (*models.Item, error)
	Delete(id uint) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}
func (r *itemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return &items, nil
}

func (r *itemRepository) FindById(id uint) (*models.Item, error) {
	var item models.Item
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) Create(item models.Item) (*models.Item, error) {
	if err := r.db.Create(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) Update(item models.Item) (*models.Item, error) {
	if err := r.db.Save(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) Delete(id uint) error {
	deleteItem, err := r.FindById(id)
	if err != nil {
		return err
	}
	if err := r.db.Delete(deleteItem).Error; err != nil {
		return err
	}
	return nil
}
