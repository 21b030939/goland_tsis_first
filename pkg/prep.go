package pkg

type Artist struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	SureName  string `json:"sure_name"`
	BandName  string `json:"band_name"`
}

type Song struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
}

func PrepareResponseSongs() []Song {
	var songs []Song

	var song Song
	song.Id = "1"
	song.Name = "Highway To Hell"
	song.Artist = "AC/DC"
	songs = append(songs, song)

	song.Id = "2"
	song.Name = "Satisfaction"
	song.Artist = "Rolling Stones"
	songs = append(songs, song)

	song.Id = "3"
	song.Name = "St. Anger"
	song.Artist = "Metallica"
	songs = append(songs, song)
	return songs
}

func PrepareResponseArtists() []Artist {
	var artists []Artist

	var artist Artist
	artist.Id = "1"
	artist.FirstName = "Angus"
	artist.SureName = "Young"
	artist.BandName = "AC/DC"
	artists = append(artists, artist)

	artist.Id = "2"
	artist.FirstName = "Brian"
	artist.SureName = "Johnson"
	artist.BandName = "AC/DC"
	artists = append(artists, artist)

	artist.Id = "3"
	artist.FirstName = "Mick"
	artist.SureName = "Jagger"
	artist.BandName = "Rolling Stones"
	artists = append(artists, artist)

	artist.Id = "4"
	artist.FirstName = "James"
	artist.SureName = "Hetfield"
	artist.BandName = "Metallica"
	artists = append(artists, artist)

	artist.Id = "5"
	artist.FirstName = "Lars"
	artist.SureName = "Ulrich"
	artist.BandName = "Metallica"
	artists = append(artists, artist)
	return artists
}
