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
	router := mux.NewRouter()

	router.HandleFunc("/todo/create", controllers.CreateTodo).Methods(http.MethodPost)
	router.HandleFunc("/test", controllers.Test).Methods(http.MethodPost)
	router.HandleFunc("/todo/{id}", controllers.GetTodoByID).Methods(http.MethodGet)

	log.Println("API is running ")

	handler := cors.Default().Handler(router)

	http.ListenAndServe(":4000", handler)

}
