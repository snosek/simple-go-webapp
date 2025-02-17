package main

import (
	"fmt"
	"net/http"

	"4pw.snosek.pl/data"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home", app.newTemplateData(r))
}

func (app *application) productsList(w http.ResponseWriter, r *http.Request) {
	templateData := app.newTemplateData(r)
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
		return
	}
	app.sessionManager.Put(r.Context(), "name", name)
	app.render(w, r, "view", withProduct(app.newTemplateData(r), product))
}

func (app *application) productsViewPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.serverError(w, r, err)
		return
	}
	name := app.sessionManager.GetString(r.Context(), "name")
	product, err := data.GetProductWithName(name)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	size := r.PostFormValue("size")
	msg := fmt.Sprintf("Zamówiono produkt %s o wartości %d zł!", polishName(product.Name), product.Price[size])
	app.sessionManager.Put(r.Context(), "flash", msg)
	http.Redirect(w, r, "/products/view/"+name, http.StatusSeeOther)
}
