package main

import (
	"log"
	"net/http"
	"todo/config"
	"todo/controllers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	config.Init()
	config.RedisInit()
	router := mux.NewRouter()

	router.HandleFunc("/todo/create", controllers.CreateTodo).Methods(http.MethodPost)
	router.HandleFunc("/todo/{id}", controllers.GetTodoByID).Methods(http.MethodGet)
	router.HandleFunc("/todos", controllers.GetAllTodos).Methods(http.MethodGet)
	router.HandleFunc("/todo/delete/{id}", controllers.DeleteTodo).Methods(http.MethodDelete)
	router.HandleFunc("/todo/update/{id}", controllers.UpdateTodo).Methods(http.MethodPut)
	router.HandleFunc("/deleteall", controllers.DeleteAllTodos).Methods(http.MethodDelete)
	log.Println("API is running ")

	handler := cors.AllowAll().Handler(router)

	http.ListenAndServe(":4000", handler)

}
