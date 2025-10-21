package service

import (
	"github.com/Estheraiyeola/driver-route-management/internal/user/dto"
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"github.com/Estheraiyeola/driver-route-management/internal/user/repository"
)

type UserService struct {
	userRepo    repository.UserRepository
	driverSvc   DriverService
	customerSvc CustomerService
}

func NewUserService(userRepo repository.UserRepository, driverSvc DriverService, customerSvc CustomerService) *UserService {
	return &UserService{
		userRepo:    userRepo,
		driverSvc:   driverSvc,
		customerSvc: customerSvc,
	}
}

func (s *UserService) CreateUser(request dto.CreateUserDTO) error {
	user := models.User{
		Name:  request.Name,
		Email: request.Email,
		Role:  request.Role,
	}

	savedUser, err := s.userRepo.Create(&user)
	if err != nil {
		return err
	}
	switch request.Role {
	case "driver":
		driverRequest := dto.CreateDriverDTO{
			UserID:        savedUser.ID,
			VehicleNumber: request.VehicleNumber,
			VehicleType:   request.VehicleType,
			Latitude:      request.Latitude,
			Longitude:     request.Longitude,
			LicenseNumber: request.LicenseNumber,
		}
		return s.driverSvc.CreateDriver(&driverRequest)
	case "customer":
		customerRequest := dto.CreateCustomerDTO{
			UserID:  savedUser.ID,
			Address: request.Address,
		}
		return s.customerSvc.CreateCustomer(&customerRequest)
	}

	return nil
}
