package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/tumivn/goblog/internal/server/config"
	"github.com/tumivn/goblog/internal/server/handlers"
	"github.com/tumivn/goblog/internal/storage"
)

var app config.AppConfig

func main() {
	config.LoadConfig(&app)

	e := echo.New()
	e.GET("/", handlers.Home)

	// Initialize database connection for the global variable `db`
	storage.InitDB(&app)

	e.Logger.Fatal(e.Start(":8080"))
}
