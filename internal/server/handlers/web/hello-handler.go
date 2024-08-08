package web

import (
	"github.com/labstack/echo/v4"
	"github.com/tumivn/goblog/internal/server"
	"github.com/tumivn/goblog/internal/server/views"
	"net/http"
)

type HelloHandler struct {
	server *server.Server
}

func NewHelloHandler(s *server.Server) *HelloHandler {
	return &HelloHandler{
		server: s,
	}
}

func (h *HelloHandler) Hello(c echo.Context) error {

	hi := views.HelloIndex("Hello", views.Hello("Tumi"))

	return RenderComponent(c, http.StatusOK, hi)
}
