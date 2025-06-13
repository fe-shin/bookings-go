package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fe-shin/bookings-go/internal/config"
	"github.com/fe-shin/bookings-go/internal/forms"
	"github.com/fe-shin/bookings-go/internal/models"
	"github.com/fe-shin/bookings-go/internal/render"
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
	reservation := models.Reservation{}

	data := make(map[string]any)
	data["reservation"] = reservation

	render.RenderHtmlTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form:    forms.New(nil),
		DataMap: data,
	})
}

// PostMakeReservation handles the posting of a reservation
func (repo *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}

	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]any)
		data["reservation"] = reservation

		render.RenderHtmlTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form:    form,
			DataMap: data,
		})

		return
	}
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

// PostAvailabilityAction is the post handler for availability form post
func (repo *Repository) PostAvailabilityAction(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("posted to availability, %s, %s", start, end)))
}

type availabilityJSONResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON is the get handler for AJAX GET request with JSON response
func (repo *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	jsonResponse := availabilityJSONResponse{
		OK:      true,
		Message: "Available",
	}

	response, err := json.MarshalIndent(jsonResponse, "", "  ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Favicon is the handler for the favicon file
func (repo *Repository) Favicon(w http.ResponseWriter, r *http.Request) {}
