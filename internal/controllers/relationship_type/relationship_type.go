package relationship_type

import (
	"family-tree/internal/models"
	"family-tree/internal/services/relationship_type"
	"net/http"

	"github.com/labstack/echo/v4"
)

type relationshipTypeController struct {
	service relationship_type.Service
}

type RelationshipTypeController interface {
	Create(c echo.Context) error
	GetByType(c echo.Context) error
	Update(c echo.Context) error
	List(c echo.Context) error
}

func NewRelationshipTypeController(service relationship_type.Service) RelationshipTypeController {
	return &relationshipTypeController{
		service: service,
	}
}

func (controller relationshipTypeController) Create(c echo.Context) error {
	var relationshipType models.RelationshipType

	err := c.Bind(&relationshipType)
	if err != nil {
		return err
	}

	err = controller.service.Create(&relationshipType)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, relationshipType)
}

func (controller relationshipTypeController) GetByType(c echo.Context) error {
	rltType := c.QueryParam("type")
	if rltType == "" {
		return models.ErrToGetParamRelationshipType
	}

	relationshipType, err := controller.service.GetByType(rltType)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, relationshipType)
}

func (controller relationshipTypeController) Update(c echo.Context) error {
	var relationshipType models.RelationshipType

	err := c.Bind(&relationshipType)
	if err != nil {
		// TODO create a model for this error
		return err
	}

	err = controller.service.Update(&relationshipType)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, relationshipType)
}

func (controller relationshipTypeController) List(c echo.Context) error {
	relationshipTypes, err := controller.service.List()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, relationshipTypes)
}
