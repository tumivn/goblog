package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
