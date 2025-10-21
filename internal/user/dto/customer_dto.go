package dto

type CreateCustomerDTO struct {
	UserID  uint   `json:"user_id" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type CustomerResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
