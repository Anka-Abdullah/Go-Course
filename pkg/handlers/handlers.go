package handlers

import (
	"net/http"

	"github.com/Anka-Abdullah/Go-Course/pkg/config"
	"github.com/Anka-Abdullah/Go-Course/pkg/render"
)

// Repo the repository type
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the Handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

//About is the handler for about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
