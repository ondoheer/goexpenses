package main

import (
	"log"
	"net/http"

	"./controllers"
	"./data"
	"./models"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var categories []models.Category

func main() {

	data.PopulateCategories()
	router := mux.NewRouter()

	// routes

	// Categories
	router.HandleFunc("/category", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/category/{id}", controllers.GetCategory).Methods("GET")
	router.HandleFunc("/category", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/category/{id}", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/category/{id}", controllers.DeleteCategory).Methods("DELETE")

	// Expenses
	router.HandleFunc("/expense", controllers.GetExpenses).Methods("GET")
	router.HandleFunc("/expense/{id}", controllers.GetExpense).Methods("GET")
	router.HandleFunc("/expense", controllers.CreateExpense).Methods("POST")
	router.HandleFunc("/expense/{id}", controllers.UpdateExpense).Methods("PUT")
	router.HandleFunc("/expense/{id}", controllers.DeleteExpense).Methods("DELETE")

	// Users
	router.HandleFunc("/user", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	// sanity check
	log.Fatal(http.ListenAndServe(":8000", router))
}
