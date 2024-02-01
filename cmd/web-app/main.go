package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/21b030939/tsis-one/pkg"
)

func main() {
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/health-check", pkg.HealthCheck).Methods("GET")
	router.HandleFunc("/songs", pkg.Songs).Methods("GET")
	router.HandleFunc("/artists", pkg.Artists).Methods("GET")
	router.HandleFunc("/songs/{id}", pkg.Song).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}


