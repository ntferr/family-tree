package relationship_type

import (
	"github.com/labstack/echo/v4"

	"family-tree/internal/services/relationship_type"
)

type relationshipTypeController struct {
	service relationship_type.Service
}

type RelationshipTypeController interface {
	Create(c echo.Context) error
	GetByType(c echo.Context) error
	Update(c echo.Context) error
}

func NewRelationshipTypeController(service relationship_type.Service) RelationshipTypeController {
	return &relationshipTypeController{
		service: service,
	}
}
