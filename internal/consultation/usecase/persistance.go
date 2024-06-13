package usecase

import (
	"context"
	"github.com/IIGabriel/clinic-management/internal/consultation"
	"github.com/IIGabriel/clinic-management/internal/models"
)

type persistanceConsultationsUseCase struct {
	repo consultation.Repository
}

func NewPersistanceConsultationsUseCase(repo consultation.Repository) consultation.PersistanceConsultationUsecase {
	return &persistanceConsultationsUseCase{repo}
}

func (s *persistanceConsultationsUseCase) FindBy(ctx context.Context, consultation models.Consultation) ([]models.Consultation, error) {
	return s.repo.FindBy(ctx, consultation)
}

func (s *persistanceConsultationsUseCase) Create(ctx context.Context, consultation *models.Consultation) error {
	return s.repo.Create(ctx, consultation)
}

func (s *persistanceConsultationsUseCase) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *persistanceConsultationsUseCase) Update(ctx context.Context, consultation *models.Consultation) error {
	return s.repo.Update(ctx, consultation)
}
