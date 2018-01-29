package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"./controllers"
	"./data"
	"./models"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "ondoheer"
	// password = ""
	dbname = "expenses"
)

var categories []models.Category

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
