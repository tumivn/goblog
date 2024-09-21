package web

import (
	"github.com/labstack/echo/v4"
	"github.com/tumivn/goblog/internal/domain/auth/services"
	"github.com/tumivn/goblog/internal/server"
	services2 "github.com/tumivn/goblog/internal/server/services"
	"net/http"
	"time"
)

type HomeHandler struct {
	server *server.Server
}

func NewHomeHandler(s *server.Server) *HomeHandler {
	return &HomeHandler{
		server: s,
	}
}

func (h *HomeHandler) Index(c echo.Context) error {
	user, _ := services2.GetCurrentUser(c, *h.server)
	var users, _ = services.GetUsers()
	var tick = time.Now().Format("20060102150405")

	return c.Render(http.StatusOK, "home.html", map[any]interface{}{
		"name":        "About",
		"msg":         "All about lehoangdung.blog!",
		"users":       users,
		"currentUser": user,
		"currentTime": tick,
	})
}
