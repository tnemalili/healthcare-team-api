package models

import "time"

type App struct {
	Version      string     `json:"version"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Updated      *time.Time `json:"updated"`
	ReleasedDate *time.Time `json:"released_date"`
}
