package web

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderComponent(c echo.Context, status int, cmp templ.Component) error {
	c.Response().WriteHeader(status)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
