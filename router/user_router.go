package router

import (
	"user/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/newuser", handler.CreateUser).Methods("POST")
	router.HandleFunc("/allusers", handler.AllUsers).Methods("POST")
	router.HandleFunc("/userbyid", handler.GetUserByID).Methods("POST")
	router.HandleFunc("/deletebyid", handler.DeleteUserById).Methods("POST")
	router.HandleFunc("/updatebyid", handler.UpdateById).Methods("POST")

	// shope handle
	router.HandleFunc("/shopkeeper/new",handler.NewShopkeper)

	return router
}
