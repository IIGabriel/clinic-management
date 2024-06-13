package http

import (
	"github.com/IIGabriel/clinic-management/internal/consultation"
	"github.com/IIGabriel/clinic-management/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type consultationController struct {
	usecase consultation.PersistanceConsultationUsecase
}

func NewConsultationController(u consultation.PersistanceConsultationUsecase) consultation.Controller {
	return &consultationController{usecase: u}
}

func (c *consultationController) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	var data models.Consultation
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	data.ID = id
	if err := c.usecase.Update(ctx.Request().Context(), &data); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update consultation"})
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *consultationController) Get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	consultations, err := c.usecase.FindBy(ctx.Request().Context(), models.Consultation{ID: id})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "could not get consultation"})
	}

	if len(consultations) == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "consultation not found"})
	}

	return ctx.JSON(http.StatusOK, consultations[0])
}

func (c *consultationController) GetMany(ctx echo.Context) error {
	var data models.Consultation
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	consultations, err := c.usecase.FindBy(ctx.Request().Context(), data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "could not get consultations"})
	}

	return ctx.JSON(http.StatusOK, consultations)
}

func (c *consultationController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	if err := c.usecase.Delete(ctx.Request().Context(), id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete consultation"})
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *consultationController) Create(ctx echo.Context) error {
	var data models.Consultation
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	if err := c.usecase.Create(ctx.Request().Context(), &data); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create consultation"})
	}

	return ctx.JSON(http.StatusCreated, data)
}
