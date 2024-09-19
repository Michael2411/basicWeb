package main

import (
	"GO-WEB/cmd/packages/config"
	"GO-WEB/cmd/packages/handlers"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
)

// use PAT  routing instead of calling each handler func in the main
func routesPAT(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}

// use PAT  routing instead of calling each handler func in the main
func routesCHI(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
