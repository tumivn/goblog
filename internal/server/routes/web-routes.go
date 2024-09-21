package routes

import (
	"github.com/tumivn/goblog/internal/server"
	"github.com/tumivn/goblog/internal/server/handlers/web"
)

func ConfigureWebRoutes(s *server.Server) {

	//Web (HTMX) handlers
	helloHandler := web.NewHelloHandler(s)
	wAuthHandler := web.NewAuthHandler(s)
	homeHandler := web.NewHomeHandler(s)
	wAboutHandler := web.NewAboutHandler(s)

	// Public Web
	s.Echo.GET("/hello", helloHandler.Hello)
	s.Echo.GET("/about", wAboutHandler.About)
	s.Echo.GET("/", homeHandler.Index)
	s.Echo.GET("/auth/login", wAuthHandler.GetSignIn)
	s.Echo.POST("/auth/login", wAuthHandler.PostSignIn)
	s.Echo.GET("/auth/logout", wAuthHandler.GetSignOut)
	s.Echo.GET("/auth/signup", wAuthHandler.GetSignUp)

}
