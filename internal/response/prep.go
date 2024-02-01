package response

import (
	"github.com/21b030939/tsis-one/internal/music"
)

func PrepareResponseSongs() []music.Song {
	var songs []music.Song

	var song music.Song
	song.Id = 1
	song.Name = "Highway To Hell"
	song.Artist = "AC/DC"
	songs = append(songs, song)

	song.Id = 2
	song.Name = "Satisfaction"
	song.Artist = "Rolling Stones"
	songs = append(songs, song)

	song.Id = 3
	song.Name = "St. Anger"
	song.Artist = "Metallica"
	songs = append(songs, song)
	return songs
}

func PrepareResponseArtists() []music.Artist {
	var artists []music.Artist

	var artist music.Artist
	artist.Id = 1
	artist.FirstName = "Angus"
	artist.SureName = "Young"
	artist.BandName = "AC/DC"
	artists = append(artists, artist)

	artist.Id = 2
	artist.FirstName = "Brian"
	artist.SureName = "Johnson"
	artist.BandName = "AC/DC"
	artists = append(artists, artist)

	artist.Id = 3
	artist.FirstName = "Mick"
	artist.SureName = "Jagger"
	artist.BandName = "Rolling Stones"
	artists = append(artists, artist)

	artist.Id = 4
	artist.FirstName = "James"
	artist.SureName = "Hetfield"
	artist.BandName = "Metallica"
	artists = append(artists, artist)

	artist.Id = 5
	artist.FirstName = "Lars"
	artist.SureName = "Ulrich"
	artist.BandName = "Metallica"
	artists = append(artists, artist)
	return artists
}
