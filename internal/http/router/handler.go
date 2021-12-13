package router

import (
	"family-tree/internal/controllers/health"
	"family-tree/internal/controllers/person"
	"family-tree/internal/controllers/relationship"
	"family-tree/internal/controllers/relationship_type"
	"family-tree/internal/http/middlewares"
	pDB "family-tree/internal/repository/postgres/person"
	rltDB "family-tree/internal/repository/postgres/relationship"
	rltTypeDB "family-tree/internal/repository/postgres/relationship_type"
	pService "family-tree/internal/services/person"
	rltService "family-tree/internal/services/relationship"
	rltTypeService "family-tree/internal/services/relationship_type"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter() *echo.Echo {
	e := echo.New()

	e.Pre(middlewares.Cors(), middlewares.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	relationDB := rltDB.NewRelationshipRepository()
	relationTypeDB := rltTypeDB.NewRelationshipTypeRepository()
	personDB := pDB.NewPersonRepository()

	relationService := rltService.NewService(personDB, relationDB, relationTypeDB)
	relationTypeService := rltTypeService.NewService(relationTypeDB)
	personService := pService.NewService(personDB)

	h := health.NewHealth()
	rltType := relationship_type.NewRelationshipTypeController(relationTypeService)
	personController := person.NewController(personService)
	rltController := relationship.NewController(relationService)

	router := e.Group("family-tree/v1")
	{
		router.GET("/health", h.Status)

		relationRouter := router.Group("/relationships")
		{
			relationRouter.POST("", rltController.Create)
			relationRouter.GET("", rltController.GetByID)
			relationRouter.GET("/list", rltController.List)
			relationRouter.GET("/members", rltController.GetRelationshipByName)

			relationTypeRouter := relationRouter.Group("/types")
			{
				relationTypeRouter.POST("", rltType.Create)
				relationTypeRouter.GET("", rltType.GetByType)
				relationTypeRouter.GET("/list", rltType.List)
				relationTypeRouter.PUT("", rltType.Update)
			}
		}

		personRouter := router.Group("/people")
		{
			personRouter.POST("", personController.Create)
			personRouter.GET("", personController.GetBy)
			personRouter.PUT("", personController.Update)
			personRouter.GET("/list", personController.List)
		}
	}

	return e
}
