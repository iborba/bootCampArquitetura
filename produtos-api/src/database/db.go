package database

import (
	"fmt"
	"os"
	"produtos-api/src/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupDatabase inicializa a conexão com o banco de dados real ou de testes
func SetupDatabase() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Verifique se estamos em ambiente de testes
	if isTesting() {
		// Em ambiente de testes, usamos a base de dados temporária
		db, err = SetupTestDatabase()

		if err != nil {
			return nil, fmt.Errorf("erro ao conectar ao banco de dados de teste: %v", err)
		}

		return db, nil
	}

	// Em produção, usamos a base de dados real
	db, err = gorm.Open(sqlite.Open("products.sqlite"), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}

	// Migrar o modelo de produto
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, fmt.Errorf("erro ao migrar o modelo de produto: %v", err)
	}

	return db, nil
}

// isTesting verifica se estamos em ambiente de testes
func isTesting() bool {
	// Verifique se a variável de ambiente TEST_ENV está definida como "true"
	return os.Getenv("TEST_ENV") == "true"
}

// SetupTestDatabase configura o banco de dados para o ambiente de testes
func SetupTestDatabase() (*gorm.DB, error) {
	// Aqui, criamos um banco de dados em memória para os testes
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao configurar banco de dados de testes: %v", err)
	}

	// Aqui podemos rodar as migrações, se necessário
	// db.AutoMigrate(&models.Product{}) // Exemplo de migração

	return db, nil
}
