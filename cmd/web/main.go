package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/tumivn/goblog/pkg/config"
	"github.com/tumivn/goblog/pkg/handlers"
	"github.com/tumivn/goblog/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {

	session := scs.New()
	session.Lifetime = 24 * 60 * 60                // 24 hours in seconds
	session.Cookie.Persist = true                  // persist session across browser restarts
	session.Cookie.SameSite = http.SameSiteLaxMode // same site lax mode
	session.Cookie.Secure = false                  //TODO: change to true in production

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app := config.AppConfig{
		UseCache:      true,
		TemplateCache: tc,
	}

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
