package relationship

import (
	"family-tree/internal/services/relationship"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service relationship.Service
}

type Controller interface {
	Create(c echo.Context) error
	GetByID(c echo.Context) error
	List(c echo.Context) error
}

func NewController(service relationship.Service) Controller {
	return &controller{
		service: service,
	}
}
