package models

import (
	"time"
)

type Prescription struct {
	ID               int       `json:"id"`
	PrescriptionDate time.Time `json:"prescription_date"`
	ConsultationID   int       `json:"consultation_id"`
}
