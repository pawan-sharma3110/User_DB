package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user/db"
	"user/modle"

	"github.com/gorilla/mux"
)

type Response struct {
	Id      int64
	Massage string
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser modle.User
	w.Header().Set("Content-Type", "Aplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewDecoder(r.Body).Decode(&newUser)
	insertId := db.InserUser(newUser)
	res := Response{
		Id:      insertId,
		Massage: "New user created",
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "internal server Error", http.StatusInternalServerError)
	}
}

func AllUsers(w http.ResponseWriter, r *http.Request) {

	AllUsers, err := db.Users()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "Aplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AllUsers)
}
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Aplication/json")
	w.WriteHeader(http.StatusOK)
	paramsId := mux.Vars(r)
	id := paramsId["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	user, err := db.GetUser(userId)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(user)

}
func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	paramsId := mux.Vars(r)
	id := paramsId["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = db.DeleteUserByID(userId)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User delete")
}
func UpdateById(w http.ResponseWriter, r *http.Request) {
	paramsId := mux.Vars(r)
	id := paramsId["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user modle.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Update user in database
	err = db.UpdateUser(userId, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
