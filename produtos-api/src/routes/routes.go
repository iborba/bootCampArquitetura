package routes

import (
	"produtos-api/src/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Definir rotas
	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", controllers.GetProductByID).Methods("GET")
	router.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	return router
}
