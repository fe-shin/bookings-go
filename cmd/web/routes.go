package main

import (
	"net/http"

	"github.com/fe-shin/bookings-go/pkg/config"
	"github.com/fe-shin/bookings-go/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	// router.Use(writeToConsole)
	router.Use(noSurf)
	router.Use(sessionLoad)

	router.Get("/", handlers.Repo.HomePage)
	router.Get("/about", handlers.Repo.AboutPage)
	router.Get("/contact", handlers.Repo.ContactPage)
	router.Get("/generals-quarters", handlers.Repo.GeneralsPage)
	router.Get("/majors-suite", handlers.Repo.MajorsPage)
	router.Get("/make-reservation", handlers.Repo.MakeReservationPage)
	router.Get("/search-availability", handlers.Repo.AvailabilityPage)
	router.Post("/search-availability", handlers.Repo.PostAvailabilityAction)
	router.Get("/search-availability-json", handlers.Repo.AvailabilityJSON)

	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return router
}
