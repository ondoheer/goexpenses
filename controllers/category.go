package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../data"
	"../models"
	"github.com/gorilla/mux"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(data.Categories)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	var cat models.Category

	if err == nil {
		result, err := cat.GetById(id)
		if err == nil {

			json.NewEncoder(w).Encode(result)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("{\"error\":\"Recurso no econtrado\"}"))
		}
	}

}

func CreateCategory(w http.ResponseWriter, r *http.Request) {

	var cat models.Category
	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		log.Println(err)
	}

	cat.ID = len(data.Categories) + 1

	data.Categories = append(data.Categories, cat)

	json.NewEncoder(w).Encode(data.Categories)

}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var tmpCat models.Category
	err := json.NewDecoder(r.Body).Decode(&tmpCat)
	if err != nil {
		log.Println(err)
	}

	for index, cat := range data.Categories {
		id, _ := strconv.Atoi(params["id"])
		if cat.ID == id {
			data.Categories[index].Label = tmpCat.Label
			data.Categories[index].Name = tmpCat.Name
			data.Categories[index].UserID = tmpCat.UserID
		}
	}

	json.NewEncoder(w).Encode(data.Categories)
}
func DeleteCategory(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var cat models.Category
	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		log.Println(err)
	}
	for index, cat := range data.Categories {

		id, _ := strconv.Atoi(params["id"])
		if cat.ID == id {
			deleteIndex := index
			fmt.Printf("%v: %v", cat.ID, id)
			data.Categories = append(data.Categories[:deleteIndex], data.Categories[deleteIndex+1:]...)

			break
		}

	}

	json.NewEncoder(w).Encode(data.Categories)

}
