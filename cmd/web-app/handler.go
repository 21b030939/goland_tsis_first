package main

import(
	"net/http"
	"encoding/json"
	"fmt"
	"log"

	"github.com/21b030939/tsis-one/pkg"
	"github.com/21b030939/tsis-one/internal/response"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func Songs(w http.ResponseWriter, r *http.Request) {
	log.Println("entering persons end point")
	var responseSongs pkg.ResponseSongs
	songs := response.PrepareResponseSongs()

	responseSongs.Songs = songs

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(responseSongs)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func Artists(w http.ResponseWriter, r *http.Request) {
	log.Println("entering persons end point")
	var responseArtists pkg.ResponseArtists
	artists := response.PrepareResponseArtists()

	responseArtists.Artists = artists

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(responseArtists)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}