package service

import (
	"github.com/Estheraiyeola/driver-route-management/internal/user/dto"
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"github.com/Estheraiyeola/driver-route-management/internal/user/repository"
)

type CustomerService interface {
	CreateCustomer(request *dto.CreateCustomerDTO) error
	FindAllCustomer() ([]models.Customer, error)
}

type customerService struct {
	repo repository.CustomerRepository
}

func CustomerServiceImpl(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) CreateCustomer(request *dto.CreateCustomerDTO) error {
	customer := models.Customer{
		UserID:  request.UserID,
		Address: request.Address,
	}
	return s.repo.CreateCustomer(&customer)
}

func (s *customerService) FindAllCustomer() ([]models.Customer, error) {
	return s.repo.FindAllCustomer()
}
