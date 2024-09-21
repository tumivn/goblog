package main

import (
	_ "github.com/lib/pq"
	"github.com/tumivn/goblog/internal/server"
	"github.com/tumivn/goblog/internal/server/routes"
)

func main() {
	s := server.NewServer()
	s.Init()
	routes.ConfigureAuthRoutes(s)

	//s.Echo.Static("/static", "./internal/server/static")

	//dir, _ := os.Getwd()
	//println(dir)

	s.Start(s.Config.Port)
}
