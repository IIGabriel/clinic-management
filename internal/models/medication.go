package models

type Medication struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
}
