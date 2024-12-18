package services

import (
	"testing"

	"produtos-api/src/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) CreateProduct(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) GetAllProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) GetProductByID(id uint) (*models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) GetProductByName(name string) ([]models.Product, error) {
	args := m.Called(name)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) GetProductsCount() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

func (m *MockProductRepository) UpdateProduct(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) DeleteProduct(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestServiceCreateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := NewProductService(mockRepo)

	product := &models.Product{Name: "Test Product", Price: 100.0}
	mockRepo.On("CreateProduct", product).Return(nil)

	err := productService.CreateProduct(product)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestServiceGetAllProducts(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := NewProductService(mockRepo)

	mockRepo.On("GetAllProducts").Return([]models.Product{
		{ID: 1, Name: "Product 1", Price: 100.0},
		{ID: 2, Name: "Product 2", Price: 150.0},
	}, nil)

	products, err := productService.GetAllProducts()
	assert.NoError(t, err)
	assert.Len(t, products, 2)
	mockRepo.AssertExpectations(t)
}

func TestServiceGetProductByID(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := NewProductService(mockRepo)

	product := &models.Product{ID: 1, Name: "Product 1", Price: 100.0}
	mockRepo.On("GetProductByID", uint(1)).Return(product, nil)

	result, err := productService.GetProductByID(1)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
	mockRepo.AssertExpectations(t)
}

func TestServiceGetProductByName(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := NewProductService(mockRepo)

	mockRepo.On("GetProductByName", "Product 1").Return([]models.Product{
		{ID: 1, Name: "Product 1", Price: 100.0},
	}, nil)

	products, err := productService.GetProductByName("Product 1")
	assert.NoError(t, err)
	assert.Len(t, products, 1)
	mockRepo.AssertExpectations(t)
}

func TestServiceGetProductsCount(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := NewProductService(mockRepo)

	mockRepo.On("GetProductsCount").Return(int64(2))

	count := productService.GetProductsCount()
	assert.Equal(t, int64(2), count)
	mockRepo.AssertExpectations(t)
}

func TestServiceUpdateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := NewProductService(mockRepo)

	product := &models.Product{ID: 1, Name: "Updated Product", Price: 120.0}
	mockRepo.On("UpdateProduct", product).Return(nil)

	err := productService.UpdateProduct(product)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestServiceDeleteProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := NewProductService(mockRepo)

	mockRepo.On("DeleteProduct", uint(1)).Return(nil)

	err := productService.DeleteProduct(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
