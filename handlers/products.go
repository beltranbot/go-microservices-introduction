package handlers

import (
	"log"
	"net/http"

	"github.com/beltranbot/go-microservices-introduction/data"
)

// Products struct
type Products struct {
	log *log.Logger
}

// NewProducts constructor
func NewProducts(log *log.Logger) *Products {
	return &Products{log}
}

// ServeHTTP func
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// handle update

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
