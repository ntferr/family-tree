package health

import "github.com/labstack/echo/v4"

type health struct{}

type Health interface {
	Status(e echo.Context) error
}

func NewHealth() Health {
	return &health{}
}
