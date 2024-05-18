package routes

import "github.com/labstack/echo/v4"

type apiRouter struct {}

func NewRouter() *apiRouter {
	return &apiRouter{}
}

func (r *apiRouter) RegisterRoutes() *echo.Echo {
	router := NewEcho()
	servicePath := router.Group("/api/v1")

	servicePath.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	return router
} 

func NewEcho() *echo.Echo {
	e := echo.New()

	return e
}

