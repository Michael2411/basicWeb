package handlers

import (
	models "GO-WEB/internal/Models"
	"GO-WEB/internal/config"
	"GO-WEB/internal/render"
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager

func getRoutes() http.Handler {
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

	tempCache, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("Cannot Create Template Cache")
		//return err
	}
	app.TemplateCache = tempCache
	app.UseCache = true // set true to avoid an error in the routes createTeplateCache funciton

	render.NewTemplates(&app)

	repo := NewRepo(&app)
	NewHandlers(repo)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	//no surf is used to ignore any post request without CSRF Token
	//mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(Repo.Home))
	mux.Get("/about", http.HandlerFunc(Repo.About))
	mux.Get("/kratos-room", http.HandlerFunc(Repo.Kratos))
	mux.Get("/batman-room", http.HandlerFunc(Repo.Batman))

	mux.Get("/search-availability", http.HandlerFunc(Repo.SearchAvailability))
	mux.Post("/search-availability", http.HandlerFunc(Repo.PostSearchAvailability))
	mux.Post("/search-availability-json", http.HandlerFunc(Repo.RoomAvailabilityJson))

	mux.Get("/contact", http.HandlerFunc(Repo.Contact))
	mux.Get("/makeReservation", http.HandlerFunc(Repo.MakeReservation))
	mux.Post("/makeReservation", http.HandlerFunc(Repo.PostMakeReservation))

	mux.Get("/reservation-summary", http.HandlerFunc(Repo.ReservationSummary))

	// to handle the static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// Load and Saves the session on Every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

//	. handlers level
//	./.. up tp internal folder
// ./../.. up to root level

var pathToTemplates = "./../../Templates"

func CreateTestTemplateCache() (map[string]*template.Template, error) {
	//myCache:=make(map[string]*template.Template)
	//--- you can also make map like this
	myCache := map[string]*template.Template{}
	//get all the files ending in page.tmpl from the Templates folder
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}
	//check if any layouts exists in the directory
	layouts, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	//range through these templates
	for _, page := range pages {
		// base return the last element of the page just to get the template name
		templateName := filepath.Base(page)

		//parse the template path in page and store it in a template named according to base path using New
		parsedTemplate, err := template.New(templateName).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			/* ParseGlob matches the layouts with the templates according to teh definitions in the templates themsevles what why we
			we parse the templates first */
			parsedTemplate, err = parsedTemplate.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[templateName] = parsedTemplate
	}
	return myCache, nil
}
