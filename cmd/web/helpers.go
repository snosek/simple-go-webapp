package main

import (
	"html/template"
	"net/http"

	"4pw.snosek.pl/ui"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, page string, data any) {
	ts, err := template.ParseFS(
		ui.Static,
		"html/base.html",
		"html/nav.html",
		"html/"+page+".html",
	)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
