package portal

import "github.com/labstack/echo/v4"

type PortalHandler interface {
	Login(c echo.Context) error
}

type portalHandlerImpl struct {}

func NewPortalHandler() PortalHandler {
	return &portalHandlerImpl{}
}
