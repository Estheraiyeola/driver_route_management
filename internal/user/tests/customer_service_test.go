package tests

import (
	"github.com/Estheraiyeola/driver-route-management/internal/config"
	"github.com/Estheraiyeola/driver-route-management/internal/user/dto"
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"github.com/Estheraiyeola/driver-route-management/internal/user/repository"
	"github.com/Estheraiyeola/driver-route-management/internal/user/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func setupTestDBForCustomerTests() {
	config.ConnectTestDB()

	config.DB.AutoMigrate(&models.Customer{})
}

func cleanDBForCustomerTests(tx *gorm.DB) {
	//err := config.DB.Migrator().DropTable(&models.Customer{})
	//if err != nil {
	//	return
	//}
	tx.Exec("DELETE FROM customers")

}

func TestCreateCustomer_Integration(t *testing.T) {
	db := config.DB
	setupTestDBForCustomerTests()
	defer cleanDBForCustomerTests(db)

	customerRepo := repository.NewCustomerRepository(db)
	customerService := service.CustomerServiceImpl(customerRepo)

	req := &dto.CreateCustomerDTO{
		UserID:  1,
		Address: "123 Main St",
	}
	err := customerService.CreateCustomer(req)
	assert.Nil(t, err)

	var customer models.Customer
	result := config.DB.First(&customer)
	assert.Nil(t, result.Error)
	assert.Equal(t, req.UserID, customer.UserID)
}

func TestGetAllCustomers_Integration(t *testing.T) {
	db := config.DB
	setupTestDBForCustomerTests()
	defer cleanDBForCustomerTests(db)

	customerRepo := repository.NewCustomerRepository(db)
	customerService := service.CustomerServiceImpl(customerRepo)

	req := &dto.CreateCustomerDTO{
		UserID:  1,
		Address: "123 Main St",
	}

	err := customerService.CreateCustomer(req)
	assert.Nil(t, err)

	customers, err := customerService.FindAllCustomer()
	assert.Nil(t, err)
	assert.Len(t, customers, 1)
}
