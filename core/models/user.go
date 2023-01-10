package models

type User struct {
	ID   string `json:"id"`
	Role *Role  `json:"role"`
	Type *Type  `json:"type"`
}
