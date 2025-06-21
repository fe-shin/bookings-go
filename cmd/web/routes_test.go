package main

import (
	"testing"

	"github.com/fe-shin/bookings-go/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	routes := routes(&app)

	switch routesType := routes.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Errorf("type is not *chi.mux, instead it's %T", routesType)
	}
}
