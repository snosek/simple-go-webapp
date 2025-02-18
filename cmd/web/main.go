package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"4pw.snosek.pl/ui"
	"github.com/alexedwards/scs/v2"
	"github.com/justinas/alice"
)

type application struct {
	logger         *slog.Logger
	sessionManager *scs.SessionManager
}

func main() {
	sessionManager := scs.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true
	app := &application{
		logger:         slog.New(slog.NewTextHandler(os.Stdout, nil)),
		sessionManager: sessionManager,
	}
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServerFS(ui.Static))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /products/list", app.productsList)
	mux.HandleFunc("POST /products/list", app.productsListPost)
	mux.HandleFunc("GET /products/view/{name}", app.productsView)
	mux.HandleFunc("POST /products/view/{name}", app.productsViewPost)

	middleware := alice.New(app.sessionManager.LoadAndSave, commonHeaders, app.logRequest)
	app.logger.Info("starting server on :4000")
	err := http.ListenAndServe(":4000", middleware.Then(mux))
	app.logger.Error(err.Error())
	os.Exit(1)
}
