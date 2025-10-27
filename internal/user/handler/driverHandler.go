package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"github.com/Estheraiyeola/driver-route-management/internal/user/service"
)

type DriverHandler struct {
	driverService service.DriverService
}

func NewDriverHandler(driverService service.DriverService) *DriverHandler {
	return &DriverHandler{driverService: driverService}
}

// GetAllDriversHandler handles GET /api/drivers
func (h *DriverHandler) GetAllDriversHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	drivers, err := h.driverService.GetAllDrivers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(drivers)
}

// GetDriverByIDHandler handles GET /api/drivers/{id}
func (h *DriverHandler) GetDriverByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") // ?id=1
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid driver id", http.StatusBadRequest)
		return
	}

	driver, err := h.driverService.GetDriverById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(driver)
}

// UpdateDriverHandler handles PUT /api/drivers/{id}
func (h *DriverHandler) UpdateDriverHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id") // ?id=1
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid driver id", http.StatusBadRequest)
		return
	}

	var driver models.Driver
	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.driverService.UpdateDriver(uint(id), driver); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "driver updated successfully"})
}
