package routes

import (
	"github.com/legangs/cms/internal/server"
	"github.com/legangs/cms/internal/server/handlers"
	"github.com/legangs/cms/internal/server/middlewares"
)

func ConfigureAuthRoutes(s *server.Server) {
	authHandler := handlers.NewAuthHandler(s)
	userHandler := handlers.CreateUserHandler(s)
	middleware := middlewares.CreateMiddleware(s)

	// Public routes
	s.Echo.POST("api/auth/login", authHandler.Login)
	s.Echo.POST("api/auth/register", userHandler.CreateUser)

	g := s.Echo.Group("/api/auth/users", middleware.IsAuthenticated)

	// Private routes
	g.GET("", userHandler.GetUsers)
	g.GET("/current-user", authHandler.GetMe)
	g.GET("/current-user/logout", authHandler.Logout)
}
