package models

type Provider struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SIM struct {
	SimNumber       string    `json:"sim_number"`
	PukNumber       string    `json:"puk_number"`
	ServiceProvider *Provider `json:"service_provider"`
}

type Device struct {
	ID       string    `json:"id"`
	IMEI     string    `json:"imei"`
	Sim      *SIM      `json:"sim"`
	App      *App      `json:"app"`
	Supplier *Supplier `json:"supplier"`
	Vehicle  *Vehicle  `json:"vehicle"`
}
