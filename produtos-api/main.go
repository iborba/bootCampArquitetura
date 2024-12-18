package main

import (
	"log"
	"net/http"
	"produtos-api/src/routes"
)

func main() {
	router := routes.SetupRoutes()
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
