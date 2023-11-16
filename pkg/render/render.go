package render

import (
	"github.com/tumivn/goblog/pkg/config"
	"github.com/tumivn/goblog/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds default data for all templates
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	//TODO: add default data
	return td
}

func RenderTemplate(w http.ResponseWriter, tName string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	td = AddDefaultData(td)
	t, ok := tc[tName]

	if !ok {
		log.Fatal("cannot get template from template cache")
	}
	err := t.Execute(w, td)
	if err != nil {
		log.Fatal("Unable to write the template to the browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	pagesCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return pagesCache, err
	}

	for _, page := range pages {
		tpml, err := template.ParseFiles(page, "./templates/base.layout.html")
		name := filepath.Base(page)

		if err != nil {
			return pagesCache, err
		}

		pagesCache[name] = tpml
	}
	return pagesCache, nil
}
