package intervals

import (
	"math/rand"
	"time"

	"github.com/innovate-technologies/DJ/at"
	"github.com/innovate-technologies/DJ/data"
	"github.com/innovate-technologies/DJ/itframe"
)

// PlaceIntervals places the intervals in a song slice
func PlaceIntervals(songs []data.Song) []data.Song {
	intervals := getCurrentIntervals()
	for _, interval := range intervals {
		songs = doInterval(songs, interval)
	}

	return songs
}

// SetReloads set the queue to reload when an interval stops playing
func SetReloads() {
	allIntervals := itframe.GetAllIntervals()
	a := at.GetInstance()

	for _, interval := range allIntervals {
		a.Add(at.Action{
			Event: "reloadQueue",
			Time:  interval.Start,
		})
		a.Add(at.Action{
			Event: "reloadQueue",
			Time:  interval.End,
		})
	}
}

func getCurrentIntervals() []data.Interval {
	now := time.Now()

	allIntervals := itframe.GetAllIntervals()
	currentIntervals := []data.Interval{}

	for _, interval := range allIntervals {
		if interval.Forever || (now.After(interval.Start) && now.Before(interval.End)) {
			currentIntervals = append(currentIntervals, interval)
		}
	}

	return currentIntervals
}

func doInterval(songs []data.Song, interval data.Interval) []data.Song {
	if len(interval.Songs) == 0 {
		return songs
	}

	count := 0
	orderCount := 0
	newSongs := []data.Song{}

	for _, song := range songs {
		newSongs = append(newSongs, song)
		countUp(&count, interval, song)
		if count >= interval.Every {
			count = 0
			newSongs = append(newSongs, getIntervalSongs(interval, &orderCount)...)
		}
	}

	return newSongs
}

func countUp(c *int, interval data.Interval, song data.Song) {
	if interval.IntervalMode == "songs" {
		*c++
	} else {
		*c += int(song.Duration)
	}
}

func getIntervalSongs(interval data.Interval, orderCount *int) (songs []data.Song) {
	songs = []data.Song{}

	// Set song.IgnoreSeperation = true
	for i := range interval.Songs {
		interval.Songs[i].IgnoreSeperation = true
	}

	if interval.IntervalType == "all" {
		songs = append(songs, interval.Songs...)
	}

	if interval.IntervalType == "random" {
		for i := 0; i < interval.SongsAtOnce; i++ {
			shuffle(interval.Songs)
			songs = append(songs, interval.Songs[0])
		}
	}

	if interval.IntervalType == "order" {
		for i := 0; i < interval.SongsAtOnce; i++ {
			songs = append(songs, interval.Songs[*orderCount])
			*orderCount++
			if *orderCount >= len(interval.Songs) {
				*orderCount = 0
			}
		}
	}

	return
}

func shuffle(slc []data.Song) {
	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
}
