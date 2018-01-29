package main

import (
	"encoding/json"
	"fmt"
	"log"
	"models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Category struct {
	ID     int    `json:"id,omitempty"`
	Label  string `json:"label,omitempty"`
	Name   string `json:"name,omitempty"`
	UserID int    `json:"userId,omitempty"`
}

var categories []models.Category

func GetCategories(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(categories)
}
func GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range categories {

		id, err := strconv.Atoi(params["id"])

		if err == nil {
			if item.ID == id {
				json.NewEncoder(w).Encode(item)
			}
		} else {
			json.NewEncoder(w).Encode(err)
		}

	}
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {

	var cat Category
	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		log.Println(err)
	}

	cat.ID = len(categories) + 1

	categories = append(categories, cat)

	fmt.Printf("%v", categories)

	json.NewEncoder(w).Encode(categories)

}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var tmpCat Category
	err := json.NewDecoder(r.Body).Decode(&tmpCat)
	if err != nil {
		log.Println(err)
	}

	for index, cat := range categories {
		id, _ := strconv.Atoi(params["id"])
		if cat.ID == id {
			categories[index].Label = tmpCat.Label
			categories[index].Name = tmpCat.Name
			categories[index].UserID = tmpCat.UserID
		}
	}

	json.NewEncoder(w).Encode(categories)
}
func DeleteCategory(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var cat Category
	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		log.Println(err)
	}
	for index, cat := range categories {

		id, _ := strconv.Atoi(params["id"])
		if cat.ID == id {
			deleteIndex := index
			fmt.Printf("%v: %v", cat.ID, id)
			categories = append(categories[:deleteIndex], categories[deleteIndex+1:]...)

			break
		}

	}

	json.NewEncoder(w).Encode(categories)

}

func main() {

	categories = append(categories, Category{ID: 3, Label: "Compras", Name: "compras", UserID: 1})
	categories = append(categories, Category{ID: 2, Label: "Fiesta", Name: "fiesta", UserID: 1})
	categories = append(categories, Category{ID: 1, Label: "Transporte", Name: "transporte", UserID: 1})
	categories = append(categories, Category{ID: 4, Label: "Viajes", Name: "viajes", UserID: 1})
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/category", GetCategories).Methods("GET")
	router.HandleFunc("/category/{id}", GetCategory).Methods("GET")
	router.HandleFunc("/category", CreateCategory).Methods("POST")
	router.HandleFunc("/category/{id}", UpdateCategory).Methods("PUT")
	router.HandleFunc("/category/{id}", DeleteCategory).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
