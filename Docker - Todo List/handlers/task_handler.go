package handlers

import (
	"encoding/json"
	"net/http"
	"todolist/db"
	"todolist/models"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	err := db.DB.Select(&tasks, "SELECT * FROM tasks ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, "can't get all tasks", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := db.DB.QueryRow("INSERT INTO tasks (name, description, completed, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id", task.Name, task.Description, task.Completed).Scan(&task.ID)
	if err != nil {
		http.Error(w, "can't create a task", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}
