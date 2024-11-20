package services

import (
	"produtos-api/src/database"
	"produtos-api/src/models"
	"testing"
)

func TestCreateAndGetProduct(t *testing.T) {
	db := database.SetupTestDatabase()
	database.DB = db

	product := &models.Product{
		Name:        "Test Product",
		Description: "Description",
		Price:       99.99,
		Stock:       10,
	}

	// Test CreateProduct
	err := CreateProduct(product)
	if err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	// Test GetAllProducts
	products, err := GetAllProducts()
	if err != nil || len(products) != 1 {
		t.Fatalf("Failed to retrieve products: %v", err)
	}

	if products[0].Name != "Test Product" {
		t.Errorf("Expected product name 'Test Product', got '%s'", products[0].Name)
	}
}
