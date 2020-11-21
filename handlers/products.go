package handlers

import (
	"context"
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

	product := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&product)
}

// UpdateProducts func
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // extract params from url
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to conver id", http.StatusInternalServerError)
		return
	}

	product := r.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &product)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Error updating product", http.StatusInternalServerError)
		return
	}
}

// KeyProduct struct
type KeyProduct struct{}

// MiddlewareValidateProduct func
func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler
		next.ServeHTTP(rw, r)
	})
}
