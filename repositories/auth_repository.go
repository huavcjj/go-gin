package repositories

import (
	"errors"
	"go-gin/models"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(user models.User) error
	FindUserByEmail(email string) (*models.User, error)
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

func (r *authRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
