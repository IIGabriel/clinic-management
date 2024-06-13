package consultation

import (
	"github.com/labstack/echo/v4"
)

type Controller interface {
	Update(ctx echo.Context) error
	Get(ctx echo.Context) error
	GetMany(ctx echo.Context) error
	Create(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
