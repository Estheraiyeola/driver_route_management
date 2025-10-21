package handler

import (
	"net/http"

	"github.com/Estheraiyeola/driver-route-management/internal/user/dto"
	"github.com/Estheraiyeola/driver-route-management/internal/user/service"
	"github.com/gin-gonic/gin"
)

type DriverHandler struct {
	service service.DriverService
}

func DriverHandlerImpl(s service.DriverService) *DriverHandler {
	return &DriverHandler{service: s}
}

func (h *DriverHandler) CreateDriver(c *gin.Context) {
	var request dto.CreateDriverDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateDriver(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create driver"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Driver created successfully"})
}

func (h *DriverHandler) GetAllDrivers(c *gin.Context) {
	drivers, err := h.service.GetAllDrivers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch drivers"})
		return
	}

	c.JSON(http.StatusOK, drivers)
}
