package main

import (
	_ "github.com/lib/pq"
	"github.com/tumivn/goblog/internal/server"
	"github.com/tumivn/goblog/internal/server/handlers"
	"github.com/tumivn/goblog/internal/server/routes"
)

func main() {
	s := server.NewServer()
	s.Init()
	routes.ConfigureAuthRoutes(s)

	s.Echo.GET("/", handlers.Home)

	s.Start(s.Config.Port)
}
