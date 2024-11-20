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
	mockRepo.On("CreateProduct", product).Return(nil) // Mock da criação do produto

	err := productService.CreateProduct(product)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t) // Verifica se o método foi chamado corretamente
}

func TestServiceGetAllProducts(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := NewProductService(mockRepo)

	// Mock da resposta do repositório
	mockRepo.On("GetAllProducts").Return([]models.Product{
		{ID: 1, Name: "Product 1", Price: 100.0},
		{ID: 2, Name: "Product 2", Price: 150.0},
	}, nil)

	products, err := productService.GetAllProducts()
	assert.NoError(t, err)
	assert.Len(t, products, 2) // Espera que a lista tenha 2 produtos
	mockRepo.AssertExpectations(t)
}

func TestServiceGetProductByID(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := NewProductService(mockRepo)

	product := &models.Product{ID: 1, Name: "Product 1", Price: 100.0}
	mockRepo.On("GetProductByID", uint(1)).Return(product, nil) // Mock da busca pelo ID

	result, err := productService.GetProductByID(1)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), result.ID) // Verifica se o ID do produto é o esperado
	mockRepo.AssertExpectations(t)
}

func TestServiceUpdateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := NewProductService(mockRepo)

	product := &models.Product{ID: 1, Name: "Updated Product", Price: 120.0}
	mockRepo.On("UpdateProduct", product).Return(nil) // Mock da atualização

	err := productService.UpdateProduct(product)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestServiceDeleteProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := NewProductService(mockRepo)

	mockRepo.On("DeleteProduct", uint(1)).Return(nil) // Mock da exclusão

	err := productService.DeleteProduct(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
