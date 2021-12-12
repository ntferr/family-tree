package person

import (
	"family-tree/internal/http/errors"
	"family-tree/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (pc controller) Create(c echo.Context) error {
	var person models.Person

	err := c.Bind(&person)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	err = pc.service.Create(&person)
	if err != nil {
		return errors.InternarServerError(c, err)
	}

	return c.JSON(http.StatusCreated, person)
}

func (pc controller) GetBy(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")

	person, err := pc.service.GetBy(id, name)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, person)
}

func (pc controller) Update(c echo.Context) error {
	var person models.Person

	err := c.Bind(&person)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	err = pc.service.Update(&person)
	if err != nil {
		return errors.InternarServerError(c, err)
	}

	return c.JSON(http.StatusOK, person)
}

func (pc controller) List(c echo.Context) error {
	people, err := pc.service.List()
	if err != nil {
		return errors.InternarServerError(c, err)
	}

	return c.JSON(http.StatusOK, people)
}
