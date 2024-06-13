package models

type Doctor struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CRM         string `json:"crm"`
	SpecialtyID int    `json:"specialty_id"`
}
