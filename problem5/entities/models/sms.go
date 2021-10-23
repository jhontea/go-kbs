package models

type SMS struct {
	PhoneNumber string `json:"phone_number"`
	Message     string `json:"message"`
}
