package data

// Clock is the info of a DJ Clock
type Clock struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Color    string `json:"color"`
	Username string `json:"username"`
	Start    struct {
		DayOfWeek int `json:"dayOfWeek"`
		Hour      int `json:"hour"`
		Minute    int `json:"minute"`
	} `json:"start"`
	End struct {
		DayOfWeek int `json:"dayOfWeek"`
		Hour      int `json:"hour"`
		Minute    int `json:"minute"`
	} `json:"end"`
	Tags []struct {
		Percent int    `json:"percent"`
		Tag     string `json:"tag"`
		ID      string `json:"_id"`
	} `json:"tags"`
}
