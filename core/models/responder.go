package models

import "time"

type Responder struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Surname   string     `json:"surname"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Stars     float64    `json:"stars"`
	Supplier  *Supplier  `json:"supplier"`
	Status    *Status    `json:"status"`
	Created   *time.Time `json:"created"`
	Modified  *time.Time `json:"modified"`
}
