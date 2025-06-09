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

	router.Get("/", http.HandlerFunc(handlers.Repo.HomePage))
	router.Get("/about", http.HandlerFunc(handlers.Repo.AboutPage))
	router.Get("/favicon.ico", http.HandlerFunc(handlers.Repo.Favicon))

	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return router
}
