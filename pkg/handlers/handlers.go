package handlers

import (
	"net/http"

	"github.com/fe-shin/bookings-go/pkg/config"
	"github.com/fe-shin/bookings-go/pkg/models"
	"github.com/fe-shin/bookings-go/pkg/render"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo is the repository object
var Repo *Repository

// CreateRepository creates a repository and then sets the fields inside of it
func CreateRepository(appConfig *config.AppConfig) {
	Repo = &Repository{
		App: appConfig,
	}
}

// HomePage is the handler for the index page
func (repo *Repository) HomePage(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderHtmlTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// AboutPage is the handler for the about page
func (repo *Repository) AboutPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "I'm a String Data from Template Data"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderHtmlTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// ContactPage is the handler for the contact page
func (repo *Repository) ContactPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "I'm a String Data from Template Data"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderHtmlTemplate(w, "contact.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// GeneralsPage is the handler for the generals quarters page
func (repo *Repository) GeneralsPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "I'm a String Data from Template Data"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderHtmlTemplate(w, "generals.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// MajorsPage is the handler for the majors suite page
func (repo *Repository) MajorsPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "I'm a String Data from Template Data"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderHtmlTemplate(w, "majors.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// MakeReservationPage is the handler for the reservation page
func (repo *Repository) MakeReservationPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "I'm a String Data from Template Data"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderHtmlTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// AvailabilityPage is the handler for the majors availability page
func (repo *Repository) AvailabilityPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "I'm a String Data from Template Data"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderHtmlTemplate(w, "search-availability.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Favicon is the handler for the favicon file
func (repo *Repository) Favicon(w http.ResponseWriter, r *http.Request) {}
