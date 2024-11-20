package controllers

import (
	"net/http"
	"net/http/httptest"
	"produtos-api/src/models"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
	controller := NewProductController(mockService)

	product := &models.Product{Name: "New Product", Price: 99.99}

	mockService.On("CreateProduct", product).Return(nil)

	productJSON := `{"name":"New Product","price":99.99}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(productJSON))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	controller.CreateProduct(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), "New Product")
	mockService.AssertExpectations(t)
}

func TestGetAllProductsController(t *testing.T) {
	mockService := new(MockProductService)
	controller := NewProductController(mockService)

	mockService.On("GetAllProducts").Return([]models.Product{
		{ID: 1, Name: "Product 1", Price: 100},
		{ID: 2, Name: "Product 2", Price: 150},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rr := httptest.NewRecorder()

	controller.GetAllProducts(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Product 1")
	assert.Contains(t, rr.Body.String(), "Product 2")
	mockService.AssertExpectations(t)
}

func TestGetProductByIDController(t *testing.T) {
	mockService := new(MockProductService)
	controller := NewProductController(mockService)

	product := &models.Product{ID: 1, Name: "Product 1", Price: 100}

	mockService.On("GetProductByID", uint(1)).Return(product, nil)

	r := mux.NewRouter()
	r.HandleFunc("/products/{id:[0-9]+}", controller.GetProductByID).Methods(http.MethodGet)

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Product 1")

	mockService.AssertExpectations(t)
}

func TestUpdateProductController(t *testing.T) {
	mockService := new(MockProductService)
	controller := NewProductController(mockService)

	product := &models.Product{ID: 1, Name: "Updated Product", Price: 120}

	mockService.On("UpdateProduct", product).Return(nil)

	r := mux.NewRouter()
	r.HandleFunc("/products/{id:[0-9]+}", controller.UpdateProduct).Methods(http.MethodPut)

	productJSON := `{"id":1,"name":"Updated Product","price":120}`
	req := httptest.NewRequest(http.MethodPut, "/products/1", strings.NewReader(productJSON))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	controller.UpdateProduct(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Updated Product")
	mockService.AssertExpectations(t)
}

func TestDeleteProductController(t *testing.T) {
	mockService := new(MockProductService)
	controller := NewProductController(mockService)

	mockService.On("DeleteProduct", uint(1)).Return(nil)

	r := mux.NewRouter()
	r.HandleFunc("/products/{id:[0-9]+}", controller.DeleteProduct).Methods(http.MethodDelete)

	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	controller.DeleteProduct(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	mockService.AssertExpectations(t)
}
