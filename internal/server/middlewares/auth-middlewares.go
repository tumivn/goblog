package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/legangs/cms/internal/server"
	"github.com/legangs/cms/ultilities"
)

type Middleware struct {
	Server *server.Server
}

func CreateMiddleware(s *server.Server) *Middleware {
	return &Middleware{
		Server: s,
	}
}

func (m *Middleware) IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {

	return func(e echo.Context) error {
		token, err := e.Cookie("token")
		if err != nil {
			return echo.ErrUnauthorized
		}

		_, err = ultilities.GetIssuer(token.Value, m.Server.Config.JwtSecret)
		if err != nil {
			return echo.ErrUnauthorized
		}
		return next(e)
	}
}
