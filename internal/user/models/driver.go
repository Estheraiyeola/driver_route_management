package models

type Driver struct {
	ID            uint    `gorm:"primaryKey"`
	UserID        uint    `json:"user_id"`
	LicenseNumber string  `json:"license_number"`
	VehicleNumber string  `json:"vehicle_number"`
	VehicleType   string  `json:"vehicle_type"`
	Available     bool    `json:"available"`
	Longitude     float64 `json:"longitude"`
	Latitude      float64 `json:"latitude"`
}
