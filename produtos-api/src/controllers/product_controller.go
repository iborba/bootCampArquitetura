package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"produtos-api/src/models"
	"produtos-api/src/services"

	"github.com/gorilla/mux"
)

// ProductController is a struct that defines the product controller
type ProductController struct {
	service services.ProductService
}

// NewProductController is a function that creates a new product controller
func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service: service}
}

// CreateProduct Cria um novo produto
// @Summary Cria um novo produto
// @Description Cria um novo produto
// @Tags produtos
// @Accept json
// @Produce json
// @Param product body models.Product true "Product data"
// @Success 201 {object} models.Product
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /products [post]
func (pc *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := pc.service.CreateProduct(&product); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetAllProducts Retorna todos os produtos
// @Summary Retorna todos os produtos
// @Description Retorna todos os produtos
// @Tags produtos
// @Accept json
// @Produce json
// @Param name query string false "Nome do produto"
// @Param count query string false "Contagem de produtos"
// @Success 200 {object} []models.Product
// @Failure 500 {object} string
// @Router /products [get]
func (pc *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		products, err := pc.service.GetProductByName(name)
		if err != nil {
			http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(products)
		return
	}

	count := r.URL.Query().Get("count")
	if count != "" {
		count := pc.service.GetProductsCount()
		response := map[string]int64{"count": count}

		json.NewEncoder(w).Encode(response)
		return
	}

	products, err := pc.service.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

// GetProductByID Retorna um produto pelo ID
// @Summary Retorna um produto pelo ID
// @Description Retorna um produto pelo ID
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Success 200 {object} models.Product
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /products/{id} [get]
func (pc *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	product, err := pc.service.GetProductByID(uint(id))
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// UpdateProduct Atualiza um produto
// @Summary Atualiza um produto
// @Description Atualiza um produto
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Param product body models.Product true "Product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /products/{id} [put]
func (pc *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
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

	if err := pc.service.UpdateProduct(&product); err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// DeleteProduct Deleta um produto
// @Summary Deleta um produto
// @Description Deleta um produto
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Success 204
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /products/{id} [delete]
func (pc *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := pc.service.DeleteProduct(uint(id)); err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
