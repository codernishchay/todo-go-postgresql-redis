package controllers

import (
	"encoding/json"
	"fmt"
	"log"
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

// response :  all todos
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	result := config.DB.Find(&models.Todo{})

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	json.NewEncoder(w).Encode(&result)
}

// response :  todo with particuler id
func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	result := config.DB.First(&todo, params["id"])
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	json.NewEncoder(w).Encode(&result)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	err := config.DB.Model(&todo).Where("id= ?", params["id"]).Updates(map[string]interface{}{"Todo": todo.Todo, "Description": todo.Description, "Date": todo.Date, "Priority": todo.Priority})
	if err != nil {
		log.Fatal(err)
	}

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result := config.DB.Where("id = ?", params["id"]).Delete(&models.Todo{})
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	json.NewEncoder(w).Encode(&result)
}
