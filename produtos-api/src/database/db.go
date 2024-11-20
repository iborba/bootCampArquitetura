package database

import (
	"log"
	"produtos-api/src/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected!")

	// Migração de models
	err = DB.AutoMigrate(&models.Product{})

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migrated!")
}

// SetupTestDatabase: Configura banco de dados em memória para testes
func SetupTestDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("Failed to migrate test database: %v", err)
	}

	return db
}
