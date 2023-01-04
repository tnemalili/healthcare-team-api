package models

type GenericMessage struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
