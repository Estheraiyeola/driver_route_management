package repository

import (
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	Create(user *models.User) (models.User, error)
	Delete(id int64) error
	Update(id int64, user models.User) error
	FindByID(id int64) (models.User, error)
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
func (r *userRepository) Delete(id int64) error {
	result := r.db.Delete(models.User{}, id)
	return result.Error
}

func (r *userRepository) Update(id int64, user models.User) error {
	result := r.db.Model(&models.User{}).Where("id = ?", id).Updates(user)
	return result.Error
}

func (r *userRepository) FindByID(id int64) (models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	return user, result.Error
}
