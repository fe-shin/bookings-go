package handlers

import (
	"fmt"
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

	render.RenderHtmlTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// AboutPage is the handler for the about page
func (repo *Repository) AboutPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "I'm a String Data from Template Data"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderHtmlTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// ContactPage is the handler for the contact page
func (repo *Repository) ContactPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "I'm a String Data from Template Data"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderHtmlTemplate(w, r, "contact.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// GeneralsPage is the handler for the generals quarters page
func (repo *Repository) GeneralsPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "I'm a String Data from Template Data"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderHtmlTemplate(w, r, "generals.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// MajorsPage is the handler for the majors suite page
func (repo *Repository) MajorsPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "I'm a String Data from Template Data"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderHtmlTemplate(w, r, "majors.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// MakeReservationPage is the handler for the reservation page
func (repo *Repository) MakeReservationPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "I'm a String Data from Template Data"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderHtmlTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// AvailabilityPage is the handler for the majors availability page
func (repo *Repository) AvailabilityPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "I'm a String Data from Template Data"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderHtmlTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (repo *Repository) PostAvailabilityAction(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("posted to availability, %s, %s", start, end)))
}

// Favicon is the handler for the favicon file
func (repo *Repository) Favicon(w http.ResponseWriter, r *http.Request) {}
