package repositories

import (
	"produtos-api/src/models"

	"gorm.io/gorm"
)

// ProductRepository define a interface para o repositório de produtos
type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
}

type ProductRepositoryDB struct {
	db *gorm.DB
}

// NewProductRepository cria uma nova instância do repositório real
func NewProductRepository(db *gorm.DB) *ProductRepositoryDB {
	return &ProductRepositoryDB{db}
}

func (repo *ProductRepositoryDB) CreateProduct(product *models.Product) error {
	return repo.db.Create(product).Error
}

func (repo *ProductRepositoryDB) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := repo.db.Find(&products).Error
	return products, err
}

func (repo *ProductRepositoryDB) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := repo.db.First(&product, id).Error
	return &product, err
}

func (repo *ProductRepositoryDB) UpdateProduct(product *models.Product) error {
	return repo.db.Save(product).Error
}

func (repo *ProductRepositoryDB) DeleteProduct(id uint) error {
	return repo.db.Delete(&models.Product{}, id).Error
}
