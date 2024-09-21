package routes

import (
	"github.com/tumivn/goblog/internal/server"
	"github.com/tumivn/goblog/internal/server/handlers/api"
	api2 "github.com/tumivn/goblog/internal/server/middlewares/api"
)

func ConfigureAuthRoutes(s *server.Server) {
	//APIs
	authHandler := api.NewAuthHandler(s)
	userHandler := api.CreateUserHandler(s)
	middleware := api2.CreateMiddleware(s)

	// Public apis
	s.Echo.POST("api/auth/login", authHandler.Login)
	s.Echo.POST("api/auth/register", userHandler.CreateUser)

	g := s.Echo.Group("/api/auth/users", middleware.IsAuthenticated)

	// Private apis
	g.GET("", userHandler.GetUsers)
	g.GET("/current-user", authHandler.GetMe)
	g.GET("/current-user/logout", authHandler.Logout)

}
