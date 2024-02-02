package routers

import (
	"encoding/json"
	"fmt"
	"gogorilla/model"
	"gogorilla/service"
	"log"
	"net/http"
)

func Gethealthcheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering Health Check")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is healthy")
}

func Getpersons(w http.ResponseWriter, r *http.Request) {

	log.Println("Get persons")

	persons, err := service.GetPersonsService()
	if err != nil {
		log.Fatal(err)
		return
	}

	var response model.Response

	response.Persons = persons

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonRes, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write(jsonRes)

}
