package models

type User struct {
	PhoneNumber string `json:"phone"`
}

type OTP struct {
	PhoneNumber string `json:"phone"`
	OTP         string `json:"otp"`
}