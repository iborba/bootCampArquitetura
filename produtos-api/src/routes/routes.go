package routes

import (
	"log"
	"produtos-api/src/controllers"
	"produtos-api/src/database"
	"produtos-api/src/repositories"
	"produtos-api/src/services"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	// Inicializa o banco de dados (real ou de testes, dependendo do ambiente)
	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Inicializar dependências
	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	// Cria um novo roteador
	router := mux.NewRouter()

	// Definir rotas
	router.HandleFunc("/products", productController.CreateProduct).Methods("POST")
	router.HandleFunc("/products", productController.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", productController.GetProductByID).Methods("GET")
	router.HandleFunc("/products/{id}", productController.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productController.DeleteProduct).Methods("DELETE")

	return router
}
