package main

import (
	"github.com/legangs/cms/internal/server"
	"github.com/legangs/cms/internal/server/handlers"
	"github.com/legangs/cms/internal/server/routes"
	_ "github.com/lib/pq"
)

c {
	s := server.NewServer()
	s.Init()
	routes.ConfigureAuthRoutes(s)

	s.Echo.GET("/", handlers.Home)

	s.Start(s.Config.Port)
}
