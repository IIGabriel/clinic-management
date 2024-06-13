package server

import (
	httpConsultation "github.com/IIGabriel/clinic-management/internal/consultation/delivery/http"
	"github.com/IIGabriel/clinic-management/internal/consultation/repository"
	"github.com/IIGabriel/clinic-management/internal/consultation/usecase"
	"github.com/labstack/echo/v4"
)

func (s *Server) MapControllers(app *echo.Echo) error {

	consultationRepo := repository.NewConsultationRepository()

	// UseCases
	consultationUseCase := usecase.NewPersistanceConsultationsUseCase(consultationRepo)

	// Controllers
	consultationController := httpConsultation.NewConsultationController(consultationUseCase)

	// Routes
	httpConsultation.MapConsulationRoutes(app.Group("/consultation"), consultationController)

	return nil
}
