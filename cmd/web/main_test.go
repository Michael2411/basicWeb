package main

import (
	"net/http"
	"testing"
)

func TestRun(t *testing.T) {
	err := run()
	if err != nil {
		t.Error("failed run()")
	}
}

func TestStartServer(t *testing.T) {
	// Create a new HTTP handler for testing. This handler will respond with HTTP 200 OK to all requests.
	handler := http.NewServeMux()

	//creating a temp handler to test the listen and serve
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // Respond with HTTP 200 OK
	})
	var err error
	go func() {
		err = startServer(":8081", handler)
		if err != nil {
			t.Error(err)
		}
	}()
}
