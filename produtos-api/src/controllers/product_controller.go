package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"produtos-api/src/models"
	"produtos-api/src/services"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := services.CreateProduct(&product); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := services.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	product, err := services.GetProductByID(uint(id))
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	product.ID = uint(id)

	if err := services.UpdateProduct(&product); err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteProduct(uint(id)); err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
