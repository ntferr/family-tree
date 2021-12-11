package router

import (
	"family-tree/internal/controllers/health"
	"family-tree/internal/controllers/relationship_type"
	"family-tree/internal/http/middlewares"
	rltTypeDB "family-tree/internal/repository/postgres/relationship_type"
	rltTypeService "family-tree/internal/services/relationship_type"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter() *echo.Echo {
	e := echo.New()

	e.Pre(middlewares.Cors(), middlewares.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	rltTypeDB := rltTypeDB.NewRelationshipTypeRepository()

	rltTypeService := rltTypeService.NewService(rltTypeDB)

	h := health.NewHealth()
	rltType := relationship_type.NewRelationshipTypeController(rltTypeService)

	router := e.Group("family-tree/v1")
	{
		router.GET("/health", h.Status)

		relationType := router.Group("/relationship_types")
		{
			relationType.POST("", rltType.Create)
			relationType.GET("", rltType.GetByType)
			relationType.GET("/list", rltType.List)
			relationType.PUT("", rltType.Update)
		}
	}

	return e
}
