package main

import (
	"GO-WEB/cmd/packages/config"
	handlers "GO-WEB/cmd/packages/handlers"
	"GO-WEB/cmd/packages/render"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change to true when in Prod
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	// to make session persist if browser closed
	session.Cookie.Persist = true
	// how strict on what this cookie applies to
	session.Cookie.SameSite = http.SameSiteLaxMode
	// insists cookie be encrypted only from https , set to false because localHost is http
	session.Cookie.Secure = app.InProduction

	app.Session = session
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
		Handler: routesCHI(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
