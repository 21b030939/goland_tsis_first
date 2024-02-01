package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/21b030939/tsis-one/internal/response"
	"github.com/21b030939/tsis-one/internal/music"
	"github.com/gorilla/mux"
)

func Song(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
  	params := mux.Vars(r)
	songs := response.PrepareResponseSongs()
  	for _, item := range songs {
    	if item.Id == params["id"] {
      		json.NewEncoder(w).Encode(item)
      		return
    	}
  	}
  	json.NewEncoder(w).Encode(&music.Song{})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func Songs(w http.ResponseWriter, r *http.Request) {
	log.Println("entering persons end point")
	var responseSongs ResponseSongs
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
	var responseArtists ResponseArtists
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