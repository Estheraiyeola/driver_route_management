package models

type Customer struct {
	UserID  uint   `json:"user_id"`
	Address string `json:"address"`
}
