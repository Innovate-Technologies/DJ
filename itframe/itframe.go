package itframe

import (
	"github.com/innovate-technologies/DJ/config"
	"github.com/innovate-technologies/DJ/data"
	resty "gopkg.in/resty.v0"
)

// API is the struct defining the ITFrame API
type API struct {
	Config *config.Config
	r      *resty.Client
}

// New gives a new ITFrame API instance
func New() *API {
	return &API{
		Config: config.GetConfig(),
		r:      resty.New().SetHostURL("https://itframe.innovatete.ch/"), // TO DO: make host changable
	}
}

func (a *API) getDJPath() string {
	return "/dj/" + a.Config.Username + "/" + a.Config.Internal["dj"]["key"]
}

//GetAllSongs gets all music of the user
func (a *API) GetAllSongs() []data.Song {
	response := []data.Song{}
	a.r.R().SetResult(&response).Get(a.getDJPath() + "/all-songs")

	return response
}

//GetSongInfo gets info of a specific song
func (a *API) GetSongInfo(id string) data.Song {
	response := data.Song{}
	a.r.R().SetResult(&response).Get(a.getDJPath() + "/song/" + id)

	return response
}

// GetAllSongsForTag returns all songs with a specific tag
func (a *API) GetAllSongsForTag(tag string) []data.Song {
	response := []data.Song{}
	a.r.R().SetResult(&response).Get(a.getDJPath() + "/songs-with-tag/" + tag)

	return response
}

// GetAllClocks gives the clocks for the account
func (a *API) GetAllClocks() []data.Clock {
	response := []data.Clock{}
	a.r.R().SetResult(&response).Get(a.getDJPath() + "/all-clocks")

	return response
}

// GetAllIntervals gives all intervals for an account
func (a *API) GetAllIntervals() []data.Interval {
	response := []data.Interval{}
	a.r.R().SetResult(&response).Get(a.getDJPath() + "/all-intervals")

	return response
}
