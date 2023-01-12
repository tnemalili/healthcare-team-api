package models

import "time"

type Vehicle struct {
	ID                 string     `json:"id"`
	Model              string     `json:"model"`
	Color              string     `json:"color"`
	RegistrationNumber string     `json:"registration_number"`
	Supplier           *Supplier  `json:"supplier"`
	Created            *time.Time `json:"created"`
	Modified           *time.Time `json:"modified"`
}
