package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type health struct{}

type Health interface {
	Status(e echo.Context) error
}

func NewHealth() Health {
	return &health{}
}

func (h *health) Status(e echo.Context) error {
	return e.JSON(http.StatusOK, echo.Map{
		"Message": "Service is OK",
	})
}
