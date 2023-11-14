package main

import (
	"fmt"
	"github.com/tumivn/goblog/pkg/config"
	"github.com/tumivn/goblog/pkg/handlers"
	"github.com/tumivn/goblog/pkg/render"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {

	app := config.AppConfig{}
	app.TemplateCache, _ = render.CreateTemplateCache()

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(err)
	}
}
