package repository

import (
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	Create(user *models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func UserRepositoryImpl(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *userRepository) Create(user *models.User) (models.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return *user, nil
}
