package web

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tumivn/goblog/internal/domain/auth/dtos"
	"github.com/tumivn/goblog/internal/domain/auth/services"
	"github.com/tumivn/goblog/internal/server"
	"github.com/tumivn/goblog/internal/server/views/home"
	"github.com/tumivn/goblog/ultilities"
	"net/http"
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
	//Check as if user is login
	var user *dtos.UserResponse
	token, _ := c.Cookie("token")
	println("token", token)
	errorMessage := ""
	if token != nil {
		issuer, err := ultilities.GetIssuer(token.Value, h.server.Config.JwtSecret)

		if err != nil {
			errorMessage = err.Error()
			fmt.Println("Unable to get issuer", err)
		}

		user, err = services.GetUserByEmail(issuer)
		if err != nil {
			errorMessage += " - " + err.Error()
			fmt.Println("Unable to get user", err)
		}
	}

	data := home.ViewModel{
		ErrorMessage: errorMessage,
		User:         user,
	}

	hi := home.HomePageIndex(" | Home page", user, home.HomePage(data))

	return RenderComponent(c, http.StatusOK, hi)
}
