package portal

import (
	"github.com/CHA-CHA-CHA-EP1/go-kanban-common-service/utils/handler"
	"github.com/labstack/echo/v4"
)

type loginPortalRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (p *portalHandlerImpl) Login(c echo.Context) error {
	var req loginPortalRequest

	body, err := handler.GetBody[loginPortalRequest](c, req)
	if err != nil {
		return c.String(400, err.Error())
	}

	err = handler.Validate(body)
	if err != nil {
		return c.String(400, err.Error())
	}

	return c.JSON(200, "OK")
}
