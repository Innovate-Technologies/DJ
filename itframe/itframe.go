package itframe

import (
	"github.com/innovate-technologies/DJ/config"
	"github.com/innovate-technologies/DJ/data"
	resty "gopkg.in/resty.v0"
)

// Config is the confug used for the requests
var Config = config.GetConfig()
var r = resty.New().SetHostURL("https://itframe.innovatete.ch/") // TO DO: make host changable

func getDJPath() string {
	return "/dj/" + Config.Username + "/" + Config.Internal["dj"]["key"]
}

//GetAllSongs gets all music of the user
func GetAllSongs() []data.Song {
	response := []data.Song{}
	r.R().SetResult(&response).Get(getDJPath() + "/all-songs")

	return response
}

//GetSongInfo gets info of a specific song
func GetSongInfo(id string) data.Song {
	response := data.Song{}
	r.R().SetResult(&response).Get(getDJPath() + "/song/" + id)

	return response
}

// GetAllSongsForTag returns all songs with a specific tag
func GetAllSongsForTag(tag string) []data.Song {
	response := []data.Song{}
	r.R().SetResult(&response).Get(getDJPath() + "/songs-with-tag/" + tag)

	return response
}

// GetAllClocks gives the clocks for the account
func GetAllClocks() []data.Clock {
	response := []data.Clock{}
	r.R().SetResult(response).Get(getDJPath() + "all-clocks")

	return response
}

// GetAllIntervals gives all intervals for an account
func GetAllIntervals() []data.Interval {
	response := []data.Interval{}
	r.R().SetResult(response).Get(getDJPath() + "all-intervals")

	return response
}
