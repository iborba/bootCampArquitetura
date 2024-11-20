package controllers

import (
	"net/http"
	"net/http/httptest"
	"produtos-api/src/models"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProductService is a mock of the ProductService
type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) CreateProduct(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductService) GetAllProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductService) GetProductByID(id uint) (*models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductService) UpdateProduct(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductService) DeleteProduct(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateProductController(t *testing.T) {
	mockService := new(MockProductService)
	controller := NewProductController(mockService) // Inject the mock service into the controller

	// Prepare test data
	product := &models.Product{Name: "New Product", Price: 99.99}

	// Mock service behavior
	mockService.On("CreateProduct", product).Return(nil) // Simulate a successful creation

	// Convert the product to JSON to simulate a POST request body
	productJSON := `{"name":"New Product","price":99.99}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(productJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	controller.CreateProduct(rr, req)

	// Assert the status code and response body
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), "New Product")
	mockService.AssertExpectations(t) // Verify the mock was called as expected
}

func TestGetAllProductsController(t *testing.T) {
	mockService := new(MockProductService)
	controller := NewProductController(mockService)

	// Mock service behavior
	mockService.On("GetAllProducts").Return([]models.Product{
		{ID: 1, Name: "Product 1", Price: 100},
		{ID: 2, Name: "Product 2", Price: 150},
	}, nil)

	// Create a GET request to fetch all products
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rr := httptest.NewRecorder()

	// Call the handler function
	controller.GetAllProducts(rr, req)

	// Assert the status code and response body
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Product 1")
	assert.Contains(t, rr.Body.String(), "Product 2")
	mockService.AssertExpectations(t)
}

func TestGetProductByIDController(t *testing.T) {
	mockService := new(MockProductService)
	controller := NewProductController(mockService)

	// Prepare the product data
	product := &models.Product{ID: 1, Name: "Product 1", Price: 100}

	// Mock service behavior
	mockService.On("GetProductByID", uint(1)).Return(product, nil)

	// Create a GET request with the product ID in the URL
	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	rr := httptest.NewRecorder()

	// Call the handler function
	controller.GetProductByID(rr, req)

	// Assert the status code and response body
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Product 1")
	mockService.AssertExpectations(t)
}

func TestUpdateProductController(t *testing.T) {
	mockService := new(MockProductService)
	controller := NewProductController(mockService)

	// Prepare the product data
	product := &models.Product{ID: 1, Name: "Updated Product", Price: 120}

	// Mock service behavior
	mockService.On("UpdateProduct", product).Return(nil)

	// Convert the product to JSON to simulate a PUT request body
	productJSON := `{"id":1,"name":"Updated Product","price":120}`
	req := httptest.NewRequest(http.MethodPut, "/products/1", strings.NewReader(productJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	controller.UpdateProduct(rr, req)

	// Assert the status code and response body
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Updated Product")
	mockService.AssertExpectations(t)
}

func TestDeleteProductController(t *testing.T) {
	mockService := new(MockProductService)
	controller := NewProductController(mockService)

	// Mock service behavior
	mockService.On("DeleteProduct", uint(1)).Return(nil)

	// Create a DELETE request
	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	rr := httptest.NewRecorder()

	// Call the handler function
	controller.DeleteProduct(rr, req)

	// Assert the status code and response body
	assert.Equal(t, http.StatusNoContent, rr.Code)
	mockService.AssertExpectations(t)
}
