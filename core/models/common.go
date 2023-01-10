package models

import "time"

type GenericMessage struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type Amount struct {
	ID       string  `json:"id"`
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

type Status struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Color       string     `json:"color"`
	Created     *time.Time `json:"created"`
}

type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Type struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewStatus(id int, name string, description string, color string) *Status {

	now := time.Now()

	return &Status{ID: id, Name: name, Description: description, Color: color, Created: &now}
}
