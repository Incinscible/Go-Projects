package handlers

import (
	"encoding/json"
	"net/http"
	"todolist/db"
	"todolist/models"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	err := db.DB.Select(&tasks, "SELECT * FROM tasks ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	taskID := mux.Vars(r)["task_id"]
	err := db.DB.Get(&task, "SELECT * FROM tasks WHERE id=$1", taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := db.DB.QueryRow("INSERT INTO tasks (name, description, completed, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id", task.Name, task.Description, task.Completed).Scan(&task.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	task.ID = mux.Vars(r)["task_id"]
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//task.ID, _ = strconv.Atoi(taskID)
	_, err := db.DB.Exec("UPDATE tasks SET name=$1, description=$2, completed=$3 WHERE id=$4", task.Name, task.Description, task.Completed, task.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := mux.Vars(r)["task_id"]
	_, err := db.DB.Exec("DELETE FROM tasks WHERE id=$1", taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
