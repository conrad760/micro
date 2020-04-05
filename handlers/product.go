package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/conrad760/micro/data"
	"github.com/gorilla/mux"
)

// Package Classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta

// A list of products returns in the response
// swagger:response productResponse
type productsResponseWrapper struct {
	// All products in the system
	//in: body
	Body []data.Product
}

// Products introduce a logger for logging
type Products struct {
	l *log.Logger
}

// NewProducts creates a new Products struct with a logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// swagger:route GET /products products listProducts
// Returns a list of products
// response:
// 200: productsResponse
//

// GetProducts returns a list of products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Get all Products")

	listProd := data.GetProducts()

	err := listProd.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Coffee is still brewing...", http.StatusInternalServerError)
	}
}

// swagger:route POST /products products listProducts
// Appends a product to the productList
// response:
// 200: productsResponse
//

// AddProduct adds a product to the products list
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Adding Products to Products List")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

// UpdateProduct changes a single product by id
func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Updating Product in Products List", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

// DeleteProduct removes a product from the products list
func (p Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
				http.Error(rw, "Unable to delete product", http.StatusBadRequest)
		return
	}
	
	p.l.Println("Deleting Product in Products List", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.DeleteProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

// KeyProduct is a key to the req
type KeyProduct struct{}

// MiddlewareValidateProduct validates Updates on products
func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Coffee is too hot", http.StatusBadRequest)
			return
		}

		//validate
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprintf("Coffee is too hot: try fixing this\n %s", err), http.StatusBadRequest)
			return

		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
