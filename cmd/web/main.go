package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/tumivn/goblog/internal/server"
	"github.com/tumivn/goblog/internal/server/routes"
	"html/template"
	"io"
)

// Define the template registry struct
// ref: https://medium.com/free-code-camp/how-to-setup-a-nested-html-template-in-the-go-echo-web-framework-670f16244bb4
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {
	s := server.NewServer()
	s.Init()

	s.Echo.Static("/static", "./internal/server/static")

	// Instantiate a template registry with an array of template set
	// Ref: https://gist.github.com/rand99/808e6e9702c00ce64803d94abff65678
	templates := make(map[string]*template.Template)
	templates["about.html"] = template.Must(template.ParseFiles("./internal/server/views/about/about.html", "./internal/server/views/base.html"))
	s.Echo.Renderer = &TemplateRegistry{
		templates: templates,
	}

	routes.ConfigureAuthRoutes(s)
	//dir, _ := os.Getwd()
	//println(dir)

	s.Start(s.Config.Port)
}
