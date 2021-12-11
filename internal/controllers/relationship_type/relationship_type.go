package relationship_type

import (
	"family-tree/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

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

	return nil
}
