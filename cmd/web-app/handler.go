package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/21b030939/tsis-one/pkg"
)

func GetSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	songs := pkg.PrepareResponseSongs()
	for _, item := range songs {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&pkg.Song{})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func GetSongs(w http.ResponseWriter, r *http.Request) {
	log.Println("entering persons end point")
	var responseSongs pkg.ResponseSongs
	songs := pkg.PrepareResponseSongs()

	responseSongs.Songs = songs

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(responseSongs)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func GetArtists(w http.ResponseWriter, r *http.Request) {
	log.Println("entering persons end point")
	var responseArtists pkg.ResponseArtists
	artists := pkg.PrepareResponseArtists()

	responseArtists.Artists = artists

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(responseArtists)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}
