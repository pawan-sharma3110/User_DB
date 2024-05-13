package router

import (
	"user/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/newuser", handler.CreateUser).Methods("POST")
	router.HandleFunc("/allusers", handler.AllUsers).Methods("GET")
	router.HandleFunc("/userbyid/{id}", handler.GetUserByID).Methods("GET")
	router.HandleFunc("/deletebyid/{id}", handler.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/updatebyid/{id}", handler.UpdateById).Methods("PUT")

	// shope handle
	router.HandleFunc("/shopkeeper/new", handler.NewShopkeper)

	return router
}
