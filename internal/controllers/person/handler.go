package person

import (
	"family-tree/internal/services/person"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service person.Service
}

type Controller interface {
	Create(c echo.Context) error
	GetBy(c echo.Context) error
}

func NewController(service person.Service) Controller {
	return &controller{
		service: service,
	}
}
