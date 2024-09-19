package main

import (
	"GO-WEB/cmd/packages/config"
	"GO-WEB/cmd/packages/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

// use PAT  routing instead of calling each handler func in the main
func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
