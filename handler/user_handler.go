package handler

import (
	"encoding/json"
	"net/http"
	"user/db"
	"user/modle"
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

}
func DeleteUserById(w http.ResponseWriter, r *http.Request) {

}
func UpdateById(w http.ResponseWriter, r *http.Request) {

}
