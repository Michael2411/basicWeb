package handlers

import (
	models "GO-WEB/cmd/packages/Models"
	"GO-WEB/cmd/packages/config"
	renders "GO-WEB/cmd/packages/render"
	"net/http"
)

// Repository Type
type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// Creates New Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repo to handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemp(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the About page Handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//preform some logic
	stringMaplocal := make(map[string]string)
	stringMaplocal["TEST"] = "TEST12312412"
	renders.RenderTemp(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMaplocal,
	})
}
