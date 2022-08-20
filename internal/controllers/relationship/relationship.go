package relationship

import (
	"family-tree/internal/http/errors"
	"family-tree/internal/models"
	"family-tree/internal/services/relationship"
	"net/http"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service relationship.Service
}

type Controller interface {
	Create(c echo.Context) error
	GetByID(c echo.Context) error
	List(c echo.Context) error
	GetRelationshipByName(c echo.Context) error
}

func NewController(service relationship.Service) Controller {
	return &controller{
		service: service,
	}
}

func (rc controller) Create(c echo.Context) error {
	var relationship models.Relationship

	err := c.Bind(&relationship)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	err = rc.service.Create(&relationship)
	if err != nil {
		return errors.InternarServerError(c, err)
	}

	return c.JSON(http.StatusCreated, relationship)
}

func (rc controller) GetByID(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return errors.BadRequest(c, models.ErrToGetParamID)
	}

	relationship, err := rc.service.GetByID(id)
	if err != nil {
		return errors.InternarServerError(c, err)
	}

	return c.JSON(http.StatusOK, relationship)
}

func (rc controller) List(c echo.Context) error {
	relationships, err := rc.service.List()
	if err != nil {
		return errors.InternarServerError(c, err)
	}

	return c.JSON(http.StatusOK, relationships)
}

func (rc controller) GetRelationshipByName(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return errors.NotFound(c, name)
	}

	data, err := rc.service.GetRelationByName(name)
	if err != nil {
		return errors.InternarServerError(c, err)
	}

	return c.JSON(http.StatusOK, data)
}
