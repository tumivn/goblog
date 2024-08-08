package services

import (
	"github.com/labstack/echo/v4"
	"github.com/tumivn/goblog/internal/domain/auth/dtos"
	"github.com/tumivn/goblog/internal/domain/auth/services"
	"github.com/tumivn/goblog/internal/server"
	"github.com/tumivn/goblog/ultilities"
)

func GetCurrentUser(c echo.Context, s server.Server) (*dtos.UserResponse, error) {

	token, err := c.Cookie("token")
	if err != nil {
		return nil, err
	}

	issuer, err := ultilities.GetIssuer(token.Value, s.Config.JwtSecret)
	if err != nil {
		return nil, err
	}

	user, err := services.GetUserByEmail(issuer)

	if err != nil {
		return nil, err
	}
	return user, nil
}
