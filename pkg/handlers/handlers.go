package handlers

import (
	"github.com/tumivn/goblog/pkg/render"
	"net/http"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
