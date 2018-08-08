package config

import (
	"os"

	resty "gopkg.in/resty.v0"
)

// Config contains the elements of the Cast config used for DJ
type Config struct {
	Username string                       `json:"username"`
	Name     string                       `json:"name"`
	Genre    string                       `json:"genre"`
	Hostname string                       `json:"hostname"`
	Internal map[string]map[string]string `json:"internal"`
	APIKey   string                       `json:"apikey"`
	Streams  []StreamConfig               `json:"streams"`
	DJ       struct {
		Enabled    bool `json:"enabled"`
		FadeLength int  `json:"fadeLength"`
	} `json:"DJ"`
	Input struct {
		SHOUTcast int `json:"SHOUTcast"`
	} `json:"input"`
}

type StreamConfig struct {
	Stream   string `json:"stream"`
	Password string `json:"password,omitempty"`
	Relay    string `json:"relay,omitempty"`
	Primary  bool   `json:"primary,omitempty"`
}

var r = resty.New().SetHostURL("https://itframe.innovatete.ch/") // TO DO: make host changable
var instance *Config

// GetConfig gives you the config for DJ
func GetConfig() *Config {
	if instance == nil {
		response := Config{}
		r.R().SetHeader("Content-Type", "application/json").SetBody(map[string]string{"username": os.Getenv("username"), "token": os.Getenv("ITFrameToken")}).SetResult(&response).Post("/cast/config")
		instance = &response
	}
	return instance
}
