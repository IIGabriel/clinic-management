package http

import (
	"github.com/IIGabriel/clinic-management/internal/consultation"
	"github.com/labstack/echo/v4"
)

func MapConsulationRoutes(group *echo.Group, h consultation.Controller) {
	group.PUT("/:id", h.Update)
	group.GET("/:id", h.Get)
	group.POST("", h.Create)
	group.GET("", h.GetMany)
	group.DELETE("/:id", h.Delete)
}
