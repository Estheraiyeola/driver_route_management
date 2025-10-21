package repository

import (
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"gorm.io/gorm"
)

type DriverRepository interface {
	FindAllDriver() ([]models.Driver, error)
	CreateDriver(user *models.Driver) error
	DeleteDriver(id uint) error
	UpdateDriver(id uint, user models.Driver) error
	GetDriverById(id uint) (*models.Driver, error)
}

type driverRepository struct {
	db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) DriverRepository {
	return &driverRepository{db: db}
}

func (r *driverRepository) FindAllDriver() ([]models.Driver, error) {
	var drivers []models.Driver
	result := r.db.Find(&drivers)
	return drivers, result.Error
}

func (r *driverRepository) CreateDriver(driver *models.Driver) error {
	return r.db.Create(driver).Error
}

func (r *driverRepository) DeleteDriver(id uint) error {
	return r.db.Delete(&models.Driver{}, id).Error
}

func (r *driverRepository) UpdateDriver(id uint, driver models.Driver) error {
	return r.db.Model(&models.Driver{}).Where("id = ?", id).Updates(models.Driver{
		UserID:        driver.UserID,
		LicenseNumber: driver.LicenseNumber,
		VehicleNumber: driver.VehicleNumber,
		VehicleType:   driver.VehicleType,
	}).Error
}

func (r *driverRepository) GetDriverById(id uint) (*models.Driver, error) {
	var driver models.Driver
	err := r.db.Where("id=?", id).First(&driver).Error
	if err != nil {
		return nil, err
	}
	return &driver, nil
}
