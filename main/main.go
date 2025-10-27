package main

import (
	"log"
	"net/http"

	"github.com/Estheraiyeola/driver-route-management/internal/config"
	"github.com/Estheraiyeola/driver-route-management/internal/user/handler"
	"github.com/Estheraiyeola/driver-route-management/internal/user/repository"
	"github.com/Estheraiyeola/driver-route-management/internal/user/service"
)

func main() {
	// Connect to database
	config.ConnectDB()

	// repositories
	userRepo := repository.UserRepositoryImpl(config.DB)
	driverRepo := repository.NewDriverRepository(config.DB)
	customerRepo := repository.NewCustomerRepository(config.DB)

	// services
	driverSvc := service.DriverServiceImpl(driverRepo)
	customerSvc := service.CustomerServiceImpl(customerRepo)
	userSvc := service.NewUserService(userRepo, driverSvc, customerSvc)

	// handlers
	driverHandler := handler.NewDriverHandler(driverSvc)
	customerHandler := handler.NewCustomerHandler(customerSvc)
	userHandler := handler.NewUserHandler(userSvc)

	http.HandleFunc("/api/users", userHandler.GetAllUsersHandler)
	http.HandleFunc("/api/users/create", userHandler.CreateUserHandler)

	http.HandleFunc("/api/drivers", driverHandler.GetAllDriversHandler)
	http.HandleFunc("/api/drivers/update", driverHandler.UpdateDriverHandler)
	http.HandleFunc("/api/drivers/get", driverHandler.GetDriverByIDHandler)

	http.HandleFunc("/api/customers", customerHandler.GetAllCustomersHandler)
	http.HandleFunc("/api/customers/get", customerHandler.GetCustomerByIDHandler)
	http.HandleFunc("/api/customers/update", customerHandler.UpdateCustomerHandler)

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
