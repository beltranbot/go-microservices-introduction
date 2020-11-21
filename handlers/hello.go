package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello struct
type Hello struct {
	log *log.Logger
}

// NewHello constructor
func NewHello(log *log.Logger) *Hello {
	return &Hello{log}
}

// ServeHTTP func
func (hello *Hello) ServeHTTP(readWriter http.ResponseWriter, request *http.Request) {
	hello.log.Println("Hello World")
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(readWriter, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(readWriter, "Hello %s", data)
}
