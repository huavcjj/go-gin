package repositories

import (
	"go-gin/models"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(user models.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateUser(user models.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
