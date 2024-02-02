package main

import (
	"gogorilla/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("routing")
	router := mux.NewRouter()
	router.HandleFunc("/health-check", routers.Gethealthcheck).Methods("GET")
	router.HandleFunc("/persons", routers.Getpersons).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
