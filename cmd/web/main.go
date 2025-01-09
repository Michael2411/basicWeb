package main

import (
	models "GO-WEB/internal/Models"
	"GO-WEB/internal/config"
	handlers "GO-WEB/internal/handlers"
	"GO-WEB/internal/render"
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	err = startServer(portNumber, routesCHI(&app))
	log.Fatal(err)
}

func startServer(port string, handler http.Handler) error {
	server := &http.Server{
		Addr:    port,
		Handler: handler,
	}
	return server.ListenAndServe()
}

func run() error {

	//Telling APP what kind of values we are going to store in the Session (specially for types we defined)
	gob.Register(models.Reservation{})
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
		return err
	}
	app.TemplateCache = tempCache
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	return err
}
