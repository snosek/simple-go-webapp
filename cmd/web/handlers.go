package main

import (
	"html/template"
	"log"
	"net/http"

	"4pw.snosek.pl/ui"
)

func home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFS(ui.Files, "html/base.html", "html/home.html", "html/nav.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
