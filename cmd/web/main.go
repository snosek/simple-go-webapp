package main

import (
	"log/slog"
	"net/http"
	"os"

	"4pw.snosek.pl/ui"
)

type application struct {
	logger *slog.Logger
}

func main() {
	app := &application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServerFS(ui.Static))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /products/list", app.productsList)
	mux.HandleFunc("POST /products/list", app.productsListPost)

	app.logger.Info("starting server on :4000...")
	err := http.ListenAndServe(":4000", app.logRequest(mux))
	app.logger.Error(err.Error())
	os.Exit(1)
}
