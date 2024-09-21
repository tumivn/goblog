package api

import (
	"github.com/labstack/echo/v4"
	"github.com/tumivn/goblog/internal/domain/auth/dtos"
	"github.com/tumivn/goblog/internal/domain/auth/services"
	"github.com/tumivn/goblog/internal/server"
	"github.com/tumivn/goblog/ultilities"
	"net/http"
	"time"
)

type AuthHandler struct {
	server *server.Server
}

func NewAuthHandler(s *server.Server) *AuthHandler {
	return &AuthHandler{
		server: s,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	u := dtos.LoginRequest{}

	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = u.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := services.AuthenticateUser(&u, h.server.Config.JwtSecret)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   res.Token,
		Expires: res.Expires,
	})

	return c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) GetMe(c echo.Context) error {

	token, err := c.Cookie("token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, nil)
	}

	issuer, err := ultilities.GetIssuer(token.Value, h.server.Config.JwtSecret)

	user, err := services.GetUserByEmail(issuer)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, nil)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) Logout(c echo.Context) error {
	_, err := c.Cookie("token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, nil)
	}
	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(1 * time.Hour),
		Path:    "/",
	})

	return c.JSON(
		http.StatusOK,
		&dtos.MessageResponse{
			Message: "Logout success",
		},
	)
}
