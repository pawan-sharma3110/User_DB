package main

import (
	"log"
	"net/http"
	"user/router"
)

func main() {
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8030", r))
}
