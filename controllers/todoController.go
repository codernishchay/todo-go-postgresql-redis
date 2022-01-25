package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/config"
	"todo/models"

	"github.com/gorilla/mux"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {

	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	// append the todo table
	if result := config.DB.Create(&todo); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}

func GetTodos(w http.ResponseWriter, r *http.Request) {

}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var todo models.Todo

	config.DB.First(&todo, params["id"])

	json.NewEncoder(w).Encode(&todo)
}

func Test(w http.ResponseWriter, r *http.Request) {

	var todo models.Todo
	jsonbody := json.NewDecoder(r.Body).Decode(&todo)
	fmt.Println(jsonbody)

	fmt.Println()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
