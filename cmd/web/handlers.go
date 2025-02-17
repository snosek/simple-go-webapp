package main

import (
	"fmt"
	"net/http"

	"4pw.snosek.pl/data"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home", nil)
}

func (app *application) productsList(w http.ResponseWriter, r *http.Request) {
	products, err := data.GetProducts()
	if err != nil {
		app.serverError(w, r, err)
	}
	fmt.Println(products)
	w.Write([]byte("tu bedzie lista produktow"))
}

func (app *application) productsListPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tu bedzie filtrowanie"))
}
