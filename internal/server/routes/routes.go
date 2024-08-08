package routes

import (
	"github.com/tumivn/goblog/internal/server"
	"github.com/tumivn/goblog/internal/server/handlers/api"
	"github.com/tumivn/goblog/internal/server/handlers/web"
	"github.com/tumivn/goblog/internal/server/middlewares"
)

func ConfigureAuthRoutes(s *server.Server) {
	//APIs
	authHandler := api.NewAuthHandler(s)
	userHandler := api.CreateUserHandler(s)
	middleware := middlewares.CreateMiddleware(s)

	//Web (HTMX) handlers
	helloHandler := web.NewHelloHandler(s)
	wAuthHandler := web.NewAuthHandler(s)
	homeHandler := web.NewHomeHandler(s)
	wAboutHandler := web.NewAboutHandler(s)

	// Public apis
	s.Echo.POST("api/auth/login", authHandler.Login)
	s.Echo.POST("api/auth/register", userHandler.CreateUser)

	// Public Web
	s.Echo.GET("/hello", helloHandler.Hello)
	s.Echo.GET("/about", wAboutHandler.About)
	s.Echo.GET("/", homeHandler.Index)
	s.Echo.GET("/auth/login", wAuthHandler.GetSignIn)
	s.Echo.POST("/auth/login", wAuthHandler.PostSignIn)
	s.Echo.GET("/auth/logout", wAuthHandler.GetSignOut)
	s.Echo.GET("/auth/signup", wAuthHandler.GetSignUp)

	g := s.Echo.Group("/api/auth/users", middleware.IsAuthenticated)

	// Private apis
	g.GET("", userHandler.GetUsers)
	g.GET("/current-user", authHandler.GetMe)
	g.GET("/current-user/logout", authHandler.Logout)

	//Web handler
}
