package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/fe-shin/bookings-go/internal/config"
	"github.com/fe-shin/bookings-go/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func CreateTemplates(a *config.AppConfig) {
	app = a
}

func addDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderHtmlTemplate parses and renders the html templates based on filename
func RenderHtmlTemplate(w http.ResponseWriter, r *http.Request, tmplFileName string, templateData *models.TemplateData) {
	var templates map[string]*template.Template
	var err error

	if app.UseCache {
		// get cached templates
		templates = app.TemplateCache
	} else {
		templates, err = CreateCachedTemplate()

		if err != nil {
			log.Fatal(err)
		}
	}

	template, ok := templates[tmplFileName]

	if !ok {
		log.Fatal("Couldn't find a cached version of template.")
	}

	// write the template to a buffer to see if template is correct
	templateBuffer := new(bytes.Buffer)

	templateData = addDefaultData(templateData, r)

	err = template.Execute(templateBuffer, templateData)

	if err != nil {
		log.Fatal(err)
	}

	_, err = templateBuffer.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}
}

// CreateCachedTemplate is the handler for template and layout caching
func CreateCachedTemplate() (map[string]*template.Template, error) {
	var cachedTemplates = map[string]*template.Template{}
	var err error

	pageFilePaths, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return cachedTemplates, err
	}

	for _, pageFilePath := range pageFilePaths {
		pageFilename := filepath.Base(pageFilePath)

		parsedTemplate, err := template.New(pageFilename).ParseFiles(pageFilePath)

		if err != nil {
			return cachedTemplates, err
		}

		layoutPaths, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return cachedTemplates, err
		}

		if len(layoutPaths) > 0 {
			parsedTemplate, err = parsedTemplate.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return cachedTemplates, err
			}
		}

		cachedTemplates[pageFilename] = parsedTemplate
	}

	return cachedTemplates, nil
}
