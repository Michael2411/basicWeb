package handlers

import (
	models "GO-WEB/cmd/packages/Models"
	"GO-WEB/cmd/packages/config"
	renders "GO-WEB/cmd/packages/render"
	"log"
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
	remoteIp := r.RemoteAddr
	log.Println(remoteIp)
	m.App.Session.Put(r.Context(), "RemoteIP", remoteIp)
	renders.RenderTemp(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the About page Handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//preform some logic
	stringMaplocal := make(map[string]string)
	stringMaplocal["TEST"] = "TEST12312412"
	remoteIp := m.App.Session.GetString(r.Context(), "RemoteIP")
	stringMaplocal["RemoteIP"] = remoteIp
	renders.RenderTemp(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMaplocal,
	})
}

func (m *Repository) Kratos(w http.ResponseWriter, r *http.Request) {
	//preform some logic

	renders.RenderTemp(w, "kratos.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Batman(w http.ResponseWriter, r *http.Request) {
	//preform some logic

	renders.RenderTemp(w, "batman.page.tmpl", &models.TemplateData{})
}
