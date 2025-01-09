package main

import (
	"GO-WEB/internal/config"
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routesCHI(&app)
	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Error(fmt.Sprintf("Expected type *Chi.Mux but got %T", v))
	}
}
