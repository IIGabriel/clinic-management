package consultation

import (
	"context"
	"github.com/IIGabriel/clinic-management/internal/models"
)

type Repository interface {
	FindBy(ctx context.Context, consultation models.Consultation) ([]models.Consultation, error)
	Create(ctx context.Context, consultation *models.Consultation) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, consultation *models.Consultation) error
}
