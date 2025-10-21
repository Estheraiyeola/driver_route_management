package dto

type CreateDriverDTO struct {
	UserID        uint    `json:"user_id" validate:"required"`
	VehicleNumber string  `json:"vehicle_number" validate:"required"`
	VehicleType   string  `json:"vehicle_type" validate:"required"`
	Latitude      float64 `json:"latitude" validate:"required"`
	Longitude     float64 `json:"longitude" validate:"required"`
	LicenseNumber string  `json:"license_number" validate:"required"`
}

type DriverResponseDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	LicenseNo string `json:"license_no"`
}
