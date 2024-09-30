package main

import (
	"GO-WEB/cmd/packages/config"
	"GO-WEB/cmd/packages/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// // use PAT  routing instead of calling each handler func in the main
// func routesPAT(app *config.AppConfig) http.Handler {
// 	mux := pat.New()
// 	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
// 	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

// 	return mux
// }

// use CHI  routing instead of calling each handler func in the main
func routesCHI(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/kratos-room", http.HandlerFunc(handlers.Repo.Kratos))
	mux.Get("/batman-room", http.HandlerFunc(handlers.Repo.Batman))
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
