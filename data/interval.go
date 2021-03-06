package data

import "time"

//Interval is the data of an interval in the database
type Interval struct {
	Username     string    `json:"username"`
	Songs        []Song    `json:"songs"`
	IntervalType string    `json:"intervalType"` // ["random", "ordered", "all"]
	Every        int       `json:"every"`
	IntervalMode string    `json:"intervalMode"` // ["songs", "seconds"]
	SongsAtOnce  int       `json:"songsAtOnce"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	Forever      bool      `json:"forever"`
	Days         []int     `json:"days"`
	DayStart     struct {
		Hour   int `json:"hour"`
		Minute int `json:"minute"`
	} `json:"dayStart"`
	DayEnd struct {
		Hour   int `json:"hour"`
		Minute int `json:"minute"`
	} `json:"dayEnd"`
}
