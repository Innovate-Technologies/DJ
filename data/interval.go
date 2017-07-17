package data

import "time"

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
}
