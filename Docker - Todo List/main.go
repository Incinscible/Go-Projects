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
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/task", handlers.CreateTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
