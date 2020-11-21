package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/beltranbot/go-microservices-introduction/data"
	"github.com/gorilla/mux"
)

// Products struct
type Products struct {
	log *log.Logger
}

// NewProducts constructor
func NewProducts(log *log.Logger) *Products {
	return &Products{log}
}

// GetProducts func
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST Products")

	// fetch the prodcuts from the datastore
	products := data.GetProducts()

	// serialize the list to JSON
	err := products.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// AddProduct func
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST Products")

	product := &data.Product{}
	err := product.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(product)
}

// UpdateProducts func
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // extract params from url
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to conver id", http.StatusInternalServerError)
		return
	}

	p.log.Println("Handle PUT Products", id)

	product := &data.Product{}
	err = product.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Error updating product", http.StatusInternalServerError)
		return
	}
}
