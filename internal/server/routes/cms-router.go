package routes

import (
	"github.com/legangs/cms/internal/server"
	"github.com/legangs/cms/internal/server/handlers/cms"
)

func ConfigureCmsRoutes(s *server.Server) {
	userHandler := cms.NewUserHandler(s)

	s.Echo.POST("api/cms/users", userHandler.CreateUser)
	s.Echo.POST("api/cms/login", userHandler.Login)
	s.Echo.GET("api/cms/users", userHandler.GetUsers)
	s.Echo.GET("api/cms/current-user", userHandler.GetMe)
}
