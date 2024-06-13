package models

type PrescriptionMedication struct {
	PrescriptionID int `json:"prescription_id"`
	MedicationID   int `json:"medication_id"`
}
