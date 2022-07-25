package routes

import (
	"encoding/json"
	"net/http"

	"github.com/DavidAgudeloValencia/go-api/db"
	"github.com/DavidAgudeloValencia/go-api/models"
	"github.com/gorilla/mux"
)

// Index
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	if users == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Users not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}

// Show
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

// Store
func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Updated
func PutUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var userUpdated models.User
	json.NewDecoder(r.Body).Decode(&user)

	params := mux.Vars(r)
	db.DB.First(&userUpdated, params["id"])

	if userUpdated.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&userUpdated).Update("NickName", user.NickName)
	db.DB.Model(&userUpdated).Update("Email", user.Email)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

// Delete
func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Delete(&user)

	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not deleted"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}
