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

// GetUsers returns every category from the database for the current user??
// at least it should but it's not doing so!
func GetUsers(w http.ResponseWriter, r *http.Request) {

	db := models.GetDB()
	users := []models.User{}

	if db.Find(&users).RecordNotFound() {

		json.NewEncoder(w).Encode([]models.User{})
	} else {
		json.NewEncoder(w).Encode(&users)
	}
}

// GetUser returns the specified category by route ID, so mux.vars can be used
// since it only parses the URL params
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := models.GetDB()

	var user models.User
	id, err := strconv.Atoi(params["id"])

	if err == nil {
		if db.First(&user, id).RecordNotFound() {

			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("{\"error\":\"Recurso no econtrado\"}"))

		} else {
			json.NewEncoder(w).Encode(&user)
		}
	}
}

// CreateUser creates a new record associated with the current user, it parses the
// request body for mux only parses the request URI
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	db := models.GetDB()
	db.Create(&user)

	w.WriteHeader(201)

	json.NewEncoder(w).Encode(&user)

}

// UpdateUser updates an existing resource based on the URI id
// it parses the PUT body content and then appends it to the same object as the
// URI requested one
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"server error\"}"))
	} else {

		b, _ := ioutil.ReadAll(r.Body)

		db := models.GetDB()

		var user models.User
		err := json.Unmarshal(b, &user)

		if err == nil {

			user.ID = id
			db.Save(&user)
			json.NewEncoder(w).Encode(&user)

		} else {
			w.WriteHeader(500)
			w.Write([]byte("{\"error\":\"server error\"}"))
		}
	}

}

// DeleteUser deletes a resource by it's id if the user owns it
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	db := models.GetDB()

	var user models.User
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"server error\"}"))
	} else {
		user.ID = id
		db.Delete(&user)

		w.WriteHeader(202) // deleted

		w.Write([]byte("{\"status\":\"deleted\"}"))
	}

}
