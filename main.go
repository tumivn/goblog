package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/tumivn/goblog/cmd/handlers"
	"github.com/tumivn/goblog/cmd/storage"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)

	// Initialize database connection for the global variable `db`
	storage.InitDB()

	e.Logger.Fatal(e.Start(":8080"))
}
