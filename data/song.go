package data

import "time"

// Song contains all teh info of a song
type Song struct {
	ID          string    `json:"_id"`
	Song        string    `json:"song"`
	Artist      string    `json:"artist"`
	Album       string    `json:"album"`
	Artwork     string    `json:"artwork"`
	Genre       string    `json:"genre"`
	InternalURL string    `json:"internalURL"`
	Size        int       `json:"size"`
	Available   bool      `json:"available"`
	Username    string    `json:"username"`
	BPM         float64   `json:"bpm"`
	Duration    float64   `json:"duration"`
	DateAdded   time.Time `json:"dateAdded"`
	Tags        []struct {
		ID       string `json:"_id"`
		Name     string `json:"name"`
		Color    string `json:"color"`
		Username string `json:"username"`
		V        int    `json:"__v"`
	} `json:"tags"`
	ProcessedURLS []struct {
		Bitrate int    `json:"bitrate"`
		URL     string `json:"url"`
	} `json:"processedURLS"`
	Type string `json:"type"`
}
