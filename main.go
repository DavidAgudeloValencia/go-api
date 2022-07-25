package main

import (
	// "fmt"

	"net/http"

	"github.com/DavidAgudeloValencia/go-api/db"
	"github.com/DavidAgudeloValencia/go-api/models"
	"github.com/DavidAgudeloValencia/go-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	// fmt.Println("Starting")
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", routes.HomeHandler)

	// Users routes
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.PutUsersHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	// Tasks routes
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.PostTasksHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.PutTasksHandler).Methods("PUT")
	r.HandleFunc("/tasks/{id}", routes.DeleteTasksHandler).Methods("DELETE")

	// Starting server
	http.ListenAndServe(":3000", r)
}
