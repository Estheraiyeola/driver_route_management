package service

import (
	"github.com/Estheraiyeola/driver-route-management/internal/user/dto"
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"github.com/Estheraiyeola/driver-route-management/internal/user/repository"
)

type CustomerService interface {
	CreateCustomer(request *dto.CreateCustomerDTO) error
	FindAllCustomer() ([]models.Customer, error)
	GetCustomerById(id uint) (*models.Customer, error)
	UpdateCustomer(id uint, customer models.Customer) error
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

func (s *customerService) GetCustomerById(id uint) (*models.Customer, error) {
	return s.repo.GetCustomerById(id)
}

func (s *customerService) UpdateCustomer(id uint, customer models.Customer) error {
	err := s.repo.UpdateCustomer(id, customer)
	if err != nil {
		return err
	}
	return nil
}
