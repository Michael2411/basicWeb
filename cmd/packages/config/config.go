package config

import (
	"html/template"
	"log"
)

// hold the global applications configurations
// I can put anything I need in this
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
