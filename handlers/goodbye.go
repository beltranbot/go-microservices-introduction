package handlers

import (
	"log"
	"net/http"
)

// Goodbye struct
type Goodbye struct {
	log *log.Logger
}

// NewGoodbye constructor
func NewGoodbye(log *log.Logger) *Goodbye {
	return &Goodbye{log}
}

// ServeHTTP method
func (goodbye *Goodbye) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte("Byeee"))
}
