package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/fe-shin/bookings-go/pkg/config"
	"github.com/fe-shin/bookings-go/pkg/handlers"
	"github.com/fe-shin/bookings-go/pkg/render"
)

const port string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main function runs the program
func main() {

	// change to true in PROD
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateCachedTemplate()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	// create repository and handlers by sending app config, and repository
	// set the app config for render package
	handlers.CreateRepository(&app)
	render.CreateTemplates(&app)

	// Serving the routes
	fmt.Println("Serving and Listening on", port)

	server := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
