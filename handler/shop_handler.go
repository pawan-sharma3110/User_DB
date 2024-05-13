package handler

import (
	"encoding/json"
	"net/http"
	"user/db"
	"user/modle"
)

func NewShopkeper(w http.ResponseWriter, r *http.Request) {
	var newShop modle.Shopkeeper
	w.Header().Set("Content-Type", "Aplication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewDecoder(r.Body).Decode(&newShop)
	dba := db.Conect()
	defer dba.Close()
	a, _ := db.InsertShopkeeper(newShop)
	json.NewEncoder(w).Encode(a)
}
