package routes

import (
	"encoding/json"
	"net/http"

	"github.com/DavidAgudeloValencia/go-api/db"
	"github.com/DavidAgudeloValencia/go-api/models"
	"github.com/gorilla/mux"
)

// Index
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task []models.Task
	db.DB.Find(&task)
	if task == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Task not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&task)
}

// Show
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Task not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&task)
}

// Store
func PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// Updated
func PutTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	var taskUpdated models.Task
	json.NewDecoder(r.Body).Decode(&task)

	params := mux.Vars(r)
	db.DB.First(&taskUpdated, params["id"])

	if taskUpdated.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Task not found"))
		return
	}

	db.DB.Model(&taskUpdated).Update("Tittle", task.Tittle)
	db.DB.Model(&taskUpdated).Update("Description", task.Description)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&task)
}

// Delete
func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Task not found"))
		return
	}

	db.DB.Delete(&task)

	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("task not deleted"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&task)
}
