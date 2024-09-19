package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// hold the global applications configurations
// I can put anything I need in this
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
