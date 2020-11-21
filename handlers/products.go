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

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// handle update

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST Products")

	// fetch the prodcuts from the datastore
	products := data.GetProducts()

	// serialize the list to JSON
	err := products.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST Products")

	product := &data.Product{}
	err := product.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(product)
}
