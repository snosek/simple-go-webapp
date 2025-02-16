package main

import (
	"fmt"
	"log"
	"net/http"

	models "4pw.snosek.pl/data"
	"4pw.snosek.pl/ui"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /products/list", productsList)
	mux.HandleFunc("POST /products/list", productsListPost)

	products := models.GetProducts()
	fmt.Print(products)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
