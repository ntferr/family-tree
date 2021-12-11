package router

import (
	"family-tree/internal/controllers/health"
	"family-tree/internal/controllers/person"
	"family-tree/internal/controllers/relationship_type"
	"family-tree/internal/http/middlewares"
	pDB "family-tree/internal/repository/postgres/person"
	rltTypeDB "family-tree/internal/repository/postgres/relationship_type"
	pService "family-tree/internal/services/person"
	rltTypeService "family-tree/internal/services/relationship_type"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter() *echo.Echo {
	e := echo.New()

	e.Pre(middlewares.Cors(), middlewares.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	relationTypeDB := rltTypeDB.NewRelationshipTypeRepository()
	personDB := pDB.NewPersonRepository()

	relationTypeService := rltTypeService.NewService(relationTypeDB)
	personService := pService.NewService(personDB)

	h := health.NewHealth()
	rltType := relationship_type.NewRelationshipTypeController(relationTypeService)
	personController := person.NewController(personService)

	router := e.Group("family-tree/v1")
	{
		router.GET("/health", h.Status)

		relationTypeRouter := router.Group("/relationships/types")
		{
			relationTypeRouter.POST("", rltType.Create)
			relationTypeRouter.GET("", rltType.GetByType)
			relationTypeRouter.GET("/list", rltType.List)
			relationTypeRouter.PUT("", rltType.Update)
		}

		personRouter := router.Group("/people")
		{
			personRouter.POST("", personController.Create)
		}
	}

	return e
}
