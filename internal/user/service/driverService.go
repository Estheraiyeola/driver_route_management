package service

import (
	"github.com/Estheraiyeola/driver-route-management/internal/user/dto"
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"github.com/Estheraiyeola/driver-route-management/internal/user/repository"
)

type DriverService interface {
	CreateDriver(request *dto.CreateDriverDTO) error
	GetAllDrivers() ([]models.Driver, error)
}

type driverService struct {
	repo repository.DriverRepository
}

func DriverServiceImpl(repo repository.DriverRepository) DriverService {
	return &driverService{repo: repo}
}

func (s *driverService) CreateDriver(request *dto.CreateDriverDTO) error {
	driver := models.Driver{
		UserID:        request.UserID,
		LicenseNumber: request.LicenseNumber,
		VehicleNumber: request.VehicleNumber,
		VehicleType:   request.VehicleType,
		Available:     true,
		Latitude:      request.Latitude,
		Longitude:     request.Longitude,
	}
	return s.repo.CreateDriver(&driver)
}

func (s *driverService) GetAllDrivers() ([]models.Driver, error) {
	return s.repo.FindAllDriver()
}
