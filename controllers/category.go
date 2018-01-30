package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
)

// GetCategories returns every category from the database for the current user??
// at least it should but it's not doing so!
func GetCategories(w http.ResponseWriter, r *http.Request) {

	db := models.GetDB()
	categories := []models.Category{}

	if db.Find(&categories).RecordNotFound() {

		json.NewEncoder(w).Encode([]models.Category{})
	} else {
		json.NewEncoder(w).Encode(&categories)
	}
}

// GetCategory returns the specified category by route ID, so mux.vars can be used
// since it only parses the URL params
func GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := models.GetDB()

	var cat models.Category
	id, err := strconv.Atoi(params["id"])

	if err == nil {
		if db.First(&cat, id).RecordNotFound() {

			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("{\"error\":\"Recurso no econtrado\"}"))

		} else {
			json.NewEncoder(w).Encode(&cat)
		}
	}
}

// CreateCategory creates a new record associated with the current user, it parses the
// request body for mux only parses the request URI
func CreateCategory(w http.ResponseWriter, r *http.Request) {

	var cat models.Category
	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		log.Println(err)
	}

	db := models.GetDB()
	db.Create(&cat)

	w.WriteHeader(201)

	json.NewEncoder(w).Encode(&cat)

}

// UpdateCategory updates an existing resource based on the URI id
// it parses the PUT body content and then appends it to the same object as the
// URI requested one
func UpdateCategory(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"server error\"}"))
	} else {

		b, _ := ioutil.ReadAll(r.Body)

		db := models.GetDB()

		var cat models.Category
		err := json.Unmarshal(b, &cat)

		if err == nil {

			cat.ID = id
			db.Save(&cat)
			json.NewEncoder(w).Encode(&cat)

		} else {
			w.WriteHeader(500)
			w.Write([]byte("{\"error\":\"server error\"}"))
		}
	}

}

// DeleteCategory deletes a resource by it's id if the user owns it
func DeleteCategory(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	db := models.GetDB()

	var cat models.Category
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"server error\"}"))
	} else {
		cat.ID = id
		db.Delete(&cat)

		w.WriteHeader(202) // deleted

		w.Write([]byte("{\"status\":\"deleted\"}"))
	}

}
