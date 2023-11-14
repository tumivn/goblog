package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templateCache = make(map[string]*template.Template)

//// RenderTemplate renders a template
//func RenderTemplate(w http.ResponseWriter, tmpl string) {
//	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
//	err := parsedTemplate.Execute(w, nil)
//	if err != nil {
//		fmt.Println("error parsing template:", err)
//	}
//}

func RenderTemplate(w http.ResponseWriter, tName string) {
	var tmpl *template.Template
	tmpl, inMap := templateCache[tName]
	log.Println(inMap)
	if !inMap {
		tmpl, _ = template.ParseFiles("./templates/"+tName, "./templates/base.layout.html")
		templateCache[tName] = tmpl
	} else {
		log.Println("using cached template")
	}

	err := tmpl.Execute(w, nil)

	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	pagesCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return pagesCache, err
	}
	for _, page := range pages {
		tpml, err := template.ParseFiles("./templates/"+page, "./templates/base.layout.html")

		if err != nil {
			return pagesCache, err
		}
		pagesCache[page] = tpml
	}
	return pagesCache, nil
}
