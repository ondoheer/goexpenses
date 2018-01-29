package main

import (
	"log"
	"net/http"

	"./controllers"
	"./data"
	"./models"
	"github.com/gorilla/mux"
)

var categories []models.Category

func main() {

	data.PopulateCategories()
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/category", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/category/{id}", controllers.GetCategory).Methods("GET")
	router.HandleFunc("/category", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/category/{id}", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/category/{id}", controllers.DeleteCategory).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
