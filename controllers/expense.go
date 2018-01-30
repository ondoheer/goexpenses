package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
)

// GetExpense returns every category from the database for the current user??
// at least it should but it's not doing so!
func GetExpenses(w http.ResponseWriter, r *http.Request) {

	db := models.GetDB()
	expenses := []models.Expense{}

	if db.Find(&expenses).RecordNotFound() {

		json.NewEncoder(w).Encode([]models.Expense{})
	} else {
		json.NewEncoder(w).Encode(&expenses)
	}
}

// GetExpense returns the specified category by route ID, so mux.vars can be used
// since it only parses the URL params
func GetExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := models.GetDB()

	var expense models.Expense
	id, err := strconv.Atoi(params["id"])

	if err == nil {
		if db.First(&expense, id).RecordNotFound() {

			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("{\"error\":\"Recurso no econtrado\"}"))

		} else {
			json.NewEncoder(w).Encode(&expense)
		}
	}
}

// CreateExpense creates a new record associated with the current user, it parses the
// request body for mux only parses the request URI
func CreateExpense(w http.ResponseWriter, r *http.Request) {

	var expense models.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		log.Println(err)
	}

	db := models.GetDB()
	dbErr := db.Create(&expense).Error

	if dbErr == nil {

		w.WriteHeader(201)

		json.NewEncoder(w).Encode(&expense)
	} else {
		msg := fmt.Sprintf(" %s", dbErr)
		w.WriteHeader(500)
		w.Write([]byte("{\"error\": \"" + msg + "\"}"))
	}

}

// UpdateExpense updates an existing resource based on the URI id
// it parses the PUT body content and then appends it to the same object as the
// URI requested one
func UpdateExpense(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"server error\"}"))
	} else {

		b, _ := ioutil.ReadAll(r.Body)

		db := models.GetDB()

		var expense models.Expense
		err := json.Unmarshal(b, &expense)

		if err == nil {

			expense.ID = id
			dbErr := db.Save(&expense).Error

			if dbErr == nil {

				json.NewEncoder(w).Encode(&expense)
			} else {
				msg := fmt.Sprintf(" %s", dbErr)
				w.WriteHeader(500)
				w.Write([]byte("{\"error\": \"" + msg + "\"}"))
			}

		} else {
			w.WriteHeader(500)
			w.Write([]byte("{\"error\":\"server error\"}"))
		}
	}

}

// DeleteExpense deletes a resource by it's id if the user owns it
func DeleteExpense(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	db := models.GetDB()

	var expense models.Expense
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"server error\"}"))
	} else {
		expense.ID = id
		dbErr := db.Delete(&expense).Error
		if dbErr == nil {

			w.WriteHeader(202) // deleted

			w.Write([]byte("{\"status\":\"deleted\"}"))
		} else {
			msg := fmt.Sprintf(" %s", dbErr)
			w.WriteHeader(500)
			w.Write([]byte("{\"error\": \"" + msg + "\"}"))
		}
	}

}
