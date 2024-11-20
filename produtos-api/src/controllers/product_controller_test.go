package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"produtos-api/src/database"
	"produtos-api/src/models"
	"produtos-api/src/routes"
	"testing"
)

func TestCreateProductEndpoint(t *testing.T) {
	db := database.SetupTestDatabase()
	database.DB = db

	router := routes.SetupRoutes()

	product := models.Product{
		Name:        "Test Product",
		Description: "Description",
		Price:       99.99,
		Stock:       10,
	}
	body, _ := json.Marshal(product)

	req := httptest.NewRequest("POST", "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201, got %d", w.Code)
	}

	var createdProduct models.Product
	json.Unmarshal(w.Body.Bytes(), &createdProduct)

	if createdProduct.Name != product.Name {
		t.Errorf("Expected product name '%s', got '%s'", product.Name, createdProduct.Name)
	}
}
