package main

import (
	"html/template"
	"net/http"
	"time"

	"4pw.snosek.pl/data"
	"4pw.snosek.pl/ui"
)

func polishName(name string) string {
	if name == "wreath" {
		return "Wianek"
	} else if name == "flowerbox" {
		return "Flower Box"
	} else if name == "bouquet" {
		return "Bukiet"
	} else {
		return ""
	}
}

var functions template.FuncMap = template.FuncMap{
	"polishName": polishName,
}

type templateData struct {
	CurrentYear int
	Products    []data.Product
}

func newTemplateData() templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
	}
}

func (app *application) render(w http.ResponseWriter, r *http.Request, page string, data templateData) {
	ts, err := template.New(page+".html").Funcs(functions).ParseFS(
		ui.Static,
		"html/base.html",
		"html/nav.html",
		"html/"+page+".html",
	)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
