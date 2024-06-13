package repository

import (
	"context"
	"github.com/IIGabriel/clinic-management/internal/consultation"
	"github.com/IIGabriel/clinic-management/internal/models"
	"github.com/IIGabriel/clinic-management/pkg/db/mysql"
	"gorm.io/gorm"
	"time"
)

type consultationRepository struct {
	db *gorm.DB
}

func NewConsultationRepository() consultation.Repository {
	return &consultationRepository{db: mysql.Mysql()}
}

func (r *consultationRepository) FindBy(ctx context.Context, consultation models.Consultation) ([]models.Consultation, error) {
	var consultations []models.Consultation

	query := "SELECT id, consultation_date, consultation_time, patient_id, doctor_id FROM consultations WHERE 1=1"

	var args []interface{}
	if consultation.ID != 0 {
		query += " AND id = ?"
		args = append(args, consultation.ID)
	}
	if !consultation.ConsultationDate.IsZero() {
		query += " AND consultation_date = ?"
		args = append(args, consultation.ConsultationDate)
	}
	if consultation.PatientID != 0 {
		query += " AND patient_id = ?"
		args = append(args, consultation.PatientID)
	}
	if consultation.DoctorID != 0 {
		query += " AND doctor_id = ?"
		args = append(args, consultation.DoctorID)
	}

	rows, err := r.db.WithContext(ctx).Raw(query, args...).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cons models.Consultation
		var consultationDateStr, consultationTimeStr string
		if err := rows.Scan(&cons.ID, &consultationDateStr, &consultationTimeStr, &cons.PatientID, &cons.DoctorID); err != nil {
			return nil, err
		}
		cons.ConsultationDate, err = time.Parse(time.RFC3339, consultationDateStr)
		if err != nil {
			return nil, err
		}
		cons.ConsultationTime, err = time.Parse("15:04:05", consultationTimeStr)
		if err != nil {
			return nil, err
		}

		consultations = append(consultations, cons)
	}

	return consultations, nil
}

func (r *consultationRepository) Create(ctx context.Context, consultation *models.Consultation) error {
	query := "INSERT INTO consultations (consultation_date, consultation_time, patient_id, doctor_id) VALUES (?, ?, ?, ?)"
	result := r.db.WithContext(ctx).Exec(query, consultation.ConsultationDate, consultation.ConsultationTime, consultation.PatientID, consultation.DoctorID)
	return result.Error
}

func (r *consultationRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM consultations WHERE id = ?"
	result := r.db.WithContext(ctx).Exec(query, id)
	return result.Error
}

func (r *consultationRepository) Update(ctx context.Context, consultation *models.Consultation) error {
	query := "UPDATE consultations SET consultation_date = ?, consultation_time = ?, patient_id = ?, doctor_id = ? WHERE id = ?"
	result := r.db.WithContext(ctx).Exec(query, consultation.ConsultationDate, consultation.ConsultationTime, consultation.PatientID, consultation.DoctorID, consultation.ID)
	return result.Error
}
