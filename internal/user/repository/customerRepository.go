package repository

import (
	"github.com/Estheraiyeola/driver-route-management/internal/config"
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAllCustomer() ([]models.Customer, error)
	CreateCustomer(user *models.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) FindAllCustomer() ([]models.Customer, error) {
	var customers []models.Customer
	result := config.DB.Find(&customers)
	return customers, result.Error
}

func (r *customerRepository) CreateCustomer(customer *models.Customer) error {
	return config.DB.Create(customer).Error
}
