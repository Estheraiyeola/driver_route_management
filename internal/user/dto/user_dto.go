package dto

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone"`
	Role     string `json:"role" binding:"required, oneof=driver customer admin"`

	// Only used if Role == "driver"
	LicenseNumber string  `json:"license_number,omitempty"`
	VehicleNumber string  `json:"vehicle_number,omitempty"`
	VehicleType   string  `json:"vehicle_type,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`

	// Only used if Role == "customer"`
	Address string `json:"address,omitempty"`
}
