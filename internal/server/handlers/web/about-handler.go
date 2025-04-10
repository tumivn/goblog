package web

import (
	"github.com/labstack/echo/v4"
	"github.com/tumivn/goblog/internal/domain/auth/services"
	"github.com/tumivn/goblog/internal/server"
	serverServices "github.com/tumivn/goblog/internal/server/services"
	"net/http"
	"time"
)

type AboutHandler struct {
	server *server.Server
}

func NewAboutHandler(s *server.Server) *AboutHandler {
	return &AboutHandler{
		server: s,
	}
}

func (h *AboutHandler) About(c echo.Context) error {
	user, err := serverServices.GetCurrentUser(c, *h.server)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get current user"})
	}

	users, err := services.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get users"})
	}

	tick := time.Now().Format("20060102150405")

	return c.Render(http.StatusOK, "about.html", map[any]interface{}{
		"name":        "About",
		"msg":         "All about lehoangdung.blog!",
		"users":       users,
		"currentUser": user,
		"currentTime": tick,
	})
}
