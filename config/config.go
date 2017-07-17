package config

import (
	"os"
	"sync"

	resty "gopkg.in/resty.v0"
)

// Config contains the elements of the Cast config used for DJ
type Config struct {
	Username string                       `json:"username"`
	Hostname string                       `json:"hostname"`
	Internal map[string]map[string]string `json:"internal"`
	APIKey   string                       `json:"apikey"`
	Streams  []struct {
		Stream   string `json:"stream"`
		Password string `json:"password,omitempty"`
		Relay    string `json:"relay,omitempty"`
		Primary  bool   `json:"primary,omitempty"`
	} `json:"streams"`
	DJ struct {
		Enabled    bool `json:"enabled"`
		FadeLength int  `json:"fadeLength"`
	} `json:"DJ"`
}

var r = resty.New().SetHostURL("https://itframe.innovatete.ch/") // TO DO: make host changable
var instance *Config
var once sync.Once

// GetConfig gives you the config for DJ
func GetConfig() *Config {
	once.Do(func() {
		response := Config{}
		r.R().SetBody(map[string]string{"username": os.Getenv("username"), "token": os.Getenv("ITFrameToken")}).SetResult(&response).Get("/cast/config")
		instance = &response
	})
	return instance
}
