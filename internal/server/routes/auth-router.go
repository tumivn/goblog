package routes

import (
	"github.com/legangs/cms/internal/server"
	"github.com/legangs/cms/internal/server/handlers/auth"
)

func ConfigureAuthRoutes(s *server.Server) {
	userHandler := auth.NewUserHandler(s)

	s.Echo.POST("api/auth/users", userHandler.CreateUser)
	s.Echo.POST("api/auth/login", userHandler.Login)
	s.Echo.GET("api/auth/users", userHandler.GetUsers)
	s.Echo.GET("api/auth/current-user", userHandler.GetMe)
	s.Echo.GET("api/auth/current-user/logout", userHandler.Logout)
}
