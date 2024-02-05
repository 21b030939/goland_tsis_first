package pkg

type ResponseSongs struct {
	Songs []Song `json:"songs"`
}

type ResponseArtists struct {
	Artists []Artist `json:"artists"`
}
