package repository

import (
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAllCustomer() ([]models.Customer, error)
	CreateCustomer(customer *models.Customer) error
	DeleteCustomer(id uint) error
	UpdateCustomer(id uint, customer models.Customer) error
	GetCustomerById(id uint) (*models.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

// Constructor
func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

// Fetch all customers
func (r *customerRepository) FindAllCustomer() ([]models.Customer, error) {
	var customers []models.Customer
	result := r.db.Find(&customers) // <- use r.db, not config.DB
	return customers, result.Error
}

// Create a new customer
func (r *customerRepository) CreateCustomer(customer *models.Customer) error {
	return r.db.Create(customer).Error // <- use r.db
}

// Delete customer by ID
func (r *customerRepository) DeleteCustomer(id uint) error {
	return r.db.Delete(&models.Customer{}, id).Error
}

// Update customer by ID
func (r *customerRepository) UpdateCustomer(id uint, customer models.Customer) error {
	return r.db.Model(&models.Customer{}).Where("id = ?", id).Updates(customer).Error
}

// Get customer by ID
func (r *customerRepository) GetCustomerById(id uint) (*models.Customer, error) {
	var customer models.Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}
