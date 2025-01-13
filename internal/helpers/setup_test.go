package helpers

import (
	"net/http"
)

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
