package tests

import (
	"os"
	"testing"

	"github.com/Estheraiyeola/driver-route-management/internal/config"
	"github.com/Estheraiyeola/driver-route-management/internal/user/dto"
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"github.com/Estheraiyeola/driver-route-management/internal/user/repository"
	"github.com/Estheraiyeola/driver-route-management/internal/user/service"
)

var (
	db          = config.DB
	userService *service.UserService
)

func TestMain(m *testing.M) {
	// Connect to the test DB
	config.ConnectTestDB()
	db = config.DB

	// Auto-migrate all tables
	err := db.AutoMigrate(&models.User{}, &models.Driver{}, &models.Customer{})
	if err != nil {
		panic("❌ failed to migrate test database: " + err.Error())
	}

	// Initialize repositories using same DB connection
	userRepo := repository.UserRepositoryImpl(db)
	driverRepo := repository.NewDriverRepository(db)
	customerRepo := repository.NewCustomerRepository(db)

	driverSvc := service.DriverServiceImpl(driverRepo)
	customerSvc := service.CustomerServiceImpl(customerRepo)
	userService = service.NewUserService(userRepo, driverSvc, customerSvc)

	code := m.Run()

	// Clean up test DB
	db.Exec("DELETE FROM drivers")
	db.Exec("DELETE FROM customers")
	db.Exec("DELETE FROM users")

	os.Exit(code)
}

func TestCreateUser_AsDriver(t *testing.T) {
	request := dto.CreateUserDTO{
		Name:          "Test Driver",
		Email:         "driver_test@example.com",
		Role:          "driver",
		VehicleNumber: "ABC123",
		VehicleType:   "Sedan",
		Latitude:      6.5244,
		Longitude:     3.3792,
		LicenseNumber: "LIC9876",
	}

	err := userService.CreateUser(request)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var user models.User
	if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		t.Fatalf("expected user to be created, got error: %v", err)
	}

	var driver models.Driver
	if err := db.Where("user_id = ?", user.ID).First(&driver).Error; err != nil {
		t.Fatalf("expected driver to be created, got error: %v", err)
	}
}

func TestCreateUser_AsCustomer(t *testing.T) {
	request := dto.CreateUserDTO{
		Name:    "Test Customer",
		Email:   "customer_test@example.com",
		Role:    "customer",
		Address: "123 Main Street",
	}

	err := userService.CreateUser(request)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var user models.User
	if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		t.Fatalf("expected user to be created, got error: %v", err)
	}

	var customer models.Customer
	if err := db.Where("user_id = ?", user.ID).First(&customer).Error; err != nil {
		t.Fatalf("expected customer to be created, got error: %v", err)
	}
}
