package main

import (
	"log"
	"net/http"
	"todolist/db"
	"todolist/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()

	router := mux.NewRouter()
	tasksRouter := router.PathPrefix("/tasks").Subrouter()
	tasksRouter.HandleFunc("", handlers.GetTasks).Methods("GET")
	tasksRouter.HandleFunc("/{task_id}", handlers.GetTask).Methods("GET")
	tasksRouter.HandleFunc("", handlers.CreateTask).Methods("POST")
	tasksRouter.HandleFunc("/{task_id}", handlers.UpdateTask).Methods("PATCH")
	tasksRouter.HandleFunc("/{task_id}", handlers.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
