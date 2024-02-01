package music

type Artist struct {
	Id       int      `json:"id"`
	FirstName     string   `json:"first_name"`
	SureName     string   `json:"sure_name"`
	BandName     string   `json:"band_name"`
}

type Song struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Artist string  `json:"artist"`
}