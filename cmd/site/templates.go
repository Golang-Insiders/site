package main

import (
	"html/template"
	"io"

	"github.com/golang-insiders/site/internal/types"
)

type Templates struct {
	templates *template.Template
}

func newTemplate() Templates {
	return Templates{
		templates: template.Must(template.ParseGlob("public/*.html")),
	}
}

func (t *Templates) render(w io.Writer, name string, data *templateData) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type templateData struct {
	Talk      *types.Talk
	TimeZones []string
	Errors    []string
}

func newTemplateData() *templateData {
	return &templateData{
		Errors: nil,
	}
}
