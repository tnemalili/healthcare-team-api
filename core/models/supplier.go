package models

import "time"

type Supplier struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Created     *time.Time `json:"created"`
	Modified    *time.Time `json:"modified"`
}
