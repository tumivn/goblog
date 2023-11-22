package main

import (
	"github.com/labstack/echo/v4"
	"github.com/legangs/cms/internal/server/config"
	"github.com/legangs/cms/internal/server/handlers"
	"github.com/legangs/cms/internal/server/handlers/cms"
	"github.com/legangs/cms/internal/storage"
	_ "github.com/lib/pq"
)

var app config.AppConfig

func main() {
	config.LoadConfig(&app)

	e := echo.New()

	// Initialize database connection for the global variable `db`
	storage.InitDB(&app)

	e.GET("/", handlers.Home)

	//cms routes
	e.POST("api/cms/users", cms.CreateUser)

	e.Logger.Fatal(e.Start(":8080"))
}
