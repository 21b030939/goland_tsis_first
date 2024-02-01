package pkg

import (
	"github.com/21b030939/tsis-one/internal/music"
)

type ResponseSongs struct {
	Songs []music.Song `json:"songs"`
}

type ResponseArtists struct {
	Artists []music.Artist `json:"artists"`
}
