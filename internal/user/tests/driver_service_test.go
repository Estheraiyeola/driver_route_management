package tests

import (
	"github.com/Estheraiyeola/driver-route-management/internal/config"
	"github.com/Estheraiyeola/driver-route-management/internal/user/dto"
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"github.com/Estheraiyeola/driver-route-management/internal/user/repository"
	"github.com/Estheraiyeola/driver-route-management/internal/user/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupTestDBForDriverTests() {
	config.ConnectTestDB()

	// Auto migrate only models used in tests
	config.DB.AutoMigrate(&models.Driver{})
}

func cleanDBForDriverTests() {
	//err := config.DB.Migrator().DropTable(&models.Driver{})
	//if err != nil {
	//	return
	//}
}

func TestCreateDriver_Integration(t *testing.T) {
	db := config.DB
	setupTestDBForDriverTests()
	defer cleanDBForDriverTests()

	driverRepo := repository.NewDriverRepository(db)
	driverService := service.DriverServiceImpl(driverRepo)

	req := dto.CreateDriverDTO{
		UserID:        1,
		LicenseNumber: "LIC567",
		VehicleNumber: "VHX123",
		VehicleType:   "Truck",
		Latitude:      6.45,
		Longitude:     3.39,
	}

	err := driverService.CreateDriver(&req)
	assert.Nil(t, err)

	var driver models.Driver
	result := config.DB.First(&driver, "license_number = ?", "LIC567")
	assert.Nil(t, result.Error)
	assert.Equal(t, "Truck", driver.VehicleType)
}

func TestGetAllDrivers_Integration(t *testing.T) {
	db := config.DB
	setupTestDBForDriverTests()
	defer cleanDBForDriverTests()

	driverRepo := repository.NewDriverRepository(db)
	driverService := service.DriverServiceImpl(driverRepo)

	config.DB.Create(&models.Driver{
		UserID:        2,
		LicenseNumber: "TEST002",
		VehicleType:   "Sedan",
	})

	drivers, err := driverService.GetAllDrivers()
	assert.Nil(t, err)
	assert.True(t, len(drivers) >= 1)
}
