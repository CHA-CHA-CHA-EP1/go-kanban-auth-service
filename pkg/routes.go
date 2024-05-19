package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/CHA-CHA-CHA-EP1/go-kanban-auth-service/pkg/api/v1/portal"
)

type apiRouter struct {}

func NewRouter() *apiRouter {
	return &apiRouter{}
}

func (r *apiRouter) RegisterRoutes() *echo.Echo {
	router := NewEcho()
	servicePath := router.Group("/api/v1")

	apiportal := portal.NewPortalHandler()

	servicePath.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	servicePath.POST("/login", apiportal.Login)

	return router
} 

func NewEcho() *echo.Echo {
	e := echo.New()

	return e
}

