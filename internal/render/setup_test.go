package render

import (
	models "GO-WEB/internal/Models"
	"GO-WEB/internal/config"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	//Telling APP what kind of values we are going to store in the Session (specially for types we defined)
	gob.Register(models.Reservation{})
	//change to true when in Prod
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	// to make session persist if browser closed
	session.Cookie.Persist = true
	// how strict on what this cookie applies to
	session.Cookie.SameSite = http.SameSiteLaxMode
	// insists cookie be encrypted only from https , set to false because localHost is http
	session.Cookie.Secure = testApp.InProduction

	testApp.Session = session
	infoLog := log.New(os.Stdout, "Info:\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "Error:\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (writer *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (writer *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

func (writer *myWriter) WriteHeader(i int) {
}
