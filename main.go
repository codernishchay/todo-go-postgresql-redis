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
	router.HandleFunc("/todo/{id}", controllers.GetTodoByID).Methods(http.MethodGet)
	router.HandleFunc("/todos", controllers.GetAllTodos).Methods(http.MethodGet)
	router.HandleFunc("/todo/delete/{id}", controllers.DeleteTodo).Methods(http.MethodDelete)
	router.HandleFunc("/todo/update/{id}", controllers.UpdateTodo).Methods(http.MethodPut)
	log.Println("API is running ")

	handler := cors.Default().Handler(router)

	http.ListenAndServe(":4000", handler)

	// todo commandline interface

	// var dummy int = 1
	// for dummy != 0 {
	// 	fmt.Println("Enter 1 to Create a new Task ")
	// 	fmt.Println("Enter 2 to Update a existing task")
	// 	fmt.Println("Enter 3 to read a task")
	// 	fmt.Println("Enter 4 to delete a task")
	// }
	// var input int
	// fmt.Scan(&input)
	// switch input {
	// case 1:
	// 	CreateTask()
	// case 2:
	// 	UpdateTask()
	// case 3:
	// 	GeTaskByID()
	// case 4:
	// 	DeleteTask()
	// default:
	// 	fmt.Println("Please Enter correct Choice")
	// 	break

	// }
}

// func CreateTask() {
// 	var task string
// 	var due_date int
// 	var priority int
// 	fmt.Println("Enter the task")
// 	fmt.Scan(&task)
// 	fmt.Println("Enter due date")
// 	fmt.Scan(&due_date)
// 	fmt.Println("Enter priority")
// 	fmt.Scan(&priority)
// 	var todo models.Todo
// 	todo.Date = due_date
// 	todo.Priority = priority
// 	todo.Todo = task
// 	if result := config.DB.Create(&todo); result.Error != nil {
// 		fmt.Println(result.Error)
// 	}

// }
