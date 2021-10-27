package handlers

import (
	"bookings/pkg/config"
	"bookings/pkg/modules"
	"bookings/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{a}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home ...
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "ip_address", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &modules.TemplateData{})
}

// About ...
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["ip_address"] = m.App.Session.GetString(r.Context(), "ip_address")

	render.RenderTemplate(w, "about.page.tmpl", &modules.TemplateData{
		StringMap: stringMap,
	})
}
