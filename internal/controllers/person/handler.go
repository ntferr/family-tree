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
	Update(c echo.Context) error
	List(c echo.Context) error
	// TODO fazer delete; Mas precisa do delete em batch da relationship primeiro
}

func NewController(service person.Service) Controller {
	return &controller{
		service: service,
	}
}
