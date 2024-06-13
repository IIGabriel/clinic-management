package models

import (
	"time"
)

type Consultation struct {
	ID               int       `json:"id"`
	ConsultationDate time.Time `json:"consultation_date"`
	ConsultationTime time.Time `json:"consultation_time"`
	PatientID        int       `json:"patient_id"`
	DoctorID         int       `json:"doctor_id"`
}
