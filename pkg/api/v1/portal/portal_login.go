package portal

import (
	"fmt"

	"github.com/CHA-CHA-CHA-EP1/go-kanban-common-service/utils/handler"
	"github.com/labstack/echo/v4"
)

type loginPortalRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p *portalHandlerImpl) Login(c echo.Context) error {
	fmt.Println("user login")
	var req loginPortalRequest

	body, err := handler.GetBody[loginPortalRequest](c, req)

	fmt.Println(err)

	if err != nil {
		return c.String(400, err.Error())
	}

	fmt.Println(body)

	return c.String(200, body.Username)
}
