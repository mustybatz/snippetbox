package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/mustybatz/snippetbox/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet
	Form        any
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {

	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// parse the files into a template set.
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Cal ParseGlob on this template set to add any partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map, using the name of the page
		cache[name] = ts
	}

	return cache, nil
}
