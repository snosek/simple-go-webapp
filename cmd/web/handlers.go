package main

import (
	"fmt"
	"net/http"

	"4pw.snosek.pl/data"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home", newTemplateData())
}

func (app *application) productsList(w http.ResponseWriter, r *http.Request) {
	templateData := newTemplateData()
	products, err := data.GetProducts()
	if err != nil {
		app.serverError(w, r, err)
	}
	templateData.Products = products
	app.render(w, r, "list", templateData)
}

func (app *application) productsListPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tu bedzie filtrowanie"))
}

func (app *application) productsView(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	w.Write([]byte(fmt.Sprintf("podglad przedmiotu %s", polishName(name))))
}
