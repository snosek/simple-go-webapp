package main

import (
	"net/http"

	"4pw.snosek.pl/data"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home", newTemplateData())
}

func (app *application) productsList(w http.ResponseWriter, r *http.Request) {
	templateData := newTemplateData()
	templateData.Products = data.Products
	app.render(w, r, "list", templateData)
}

func (app *application) productsListPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tu bedzie filtrowanie"))
}

func (app *application) productsView(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	product, err := data.GetProductWithName(name)
	if err != nil {
		app.serverError(w, r, err)
	}
	templateData := newTemplateData()
	templateData.Product = product
	app.render(w, r, "view", templateData)
}
