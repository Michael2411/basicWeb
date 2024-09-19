package main

import (
	"GO-WEB/cmd/packages/config"
	handlers "GO-WEB/cmd/packages/handlers"
	"GO-WEB/cmd/packages/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tempCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot Create Template Cache")
	}
	app.TemplateCache = tempCache
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting App on Port %s", portNumber)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
