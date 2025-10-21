package repository

import (
	"github.com/Estheraiyeola/driver-route-management/internal/config"
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"gorm.io/gorm"
)

type DriverRepository interface {
	FindAllDriver() ([]models.Driver, error)
	CreateDriver(user *models.Driver) error
}

type driverRepository struct {
	db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) DriverRepository {
	return &driverRepository{db: db}
}

func (r *driverRepository) FindAllDriver() ([]models.Driver, error) {
	var drivers []models.Driver
	result := config.DB.Find(&drivers)
	return drivers, result.Error
}

func (r *driverRepository) CreateDriver(driver *models.Driver) error {
	return config.DB.Create(driver).Error
}
