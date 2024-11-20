package services

import (
	"produtos-api/src/models"
	"produtos-api/src/repositories"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
}

type ProductServiceRepo struct {
	repository repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductServiceRepo {
	return &ProductServiceRepo{repository: repo}
}

func (s *ProductServiceRepo) CreateProduct(product *models.Product) error {
	return s.repository.CreateProduct(product)
}

func (s *ProductServiceRepo) GetAllProducts() ([]models.Product, error) {
	return s.repository.GetAllProducts()
}

func (s *ProductServiceRepo) GetProductByID(id uint) (*models.Product, error) {
	return s.repository.GetProductByID(id)
}

func (s *ProductServiceRepo) UpdateProduct(product *models.Product) error {
	return s.repository.UpdateProduct(product)
}

func (s *ProductServiceRepo) DeleteProduct(id uint) error {
	return s.repository.DeleteProduct(id)
}
