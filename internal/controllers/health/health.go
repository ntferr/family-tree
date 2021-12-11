package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *health) Status(e echo.Context) error {
	return e.JSON(http.StatusOK, echo.Map{
		"Message": "Service is OK",
	})
}
