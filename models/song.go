package models

//Song defines a song
type Song struct {
	// Song name
	//
	// Required: true
	Name string `json:"name"`
	// Artist name
	//
	// Required: true
	Artist string `json:"artist"`
}
