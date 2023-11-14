package main

import (
	"fmt"
	"github.com/tumivn/goblog/pkg/config"
	"github.com/tumivn/goblog/pkg/handlers"
	"github.com/tumivn/goblog/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app := config.AppConfig{
		UseCache:      false,
		TemplateCache: tc,
	}
	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(err)
	}
}
