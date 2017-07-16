package config

// Config contains the elements of the Cast config used for DJ
type Config struct {
	Username string                       `json:"username"`
	Internal map[string]map[string]string `json:"internal"`
}
