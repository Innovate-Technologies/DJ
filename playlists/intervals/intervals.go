package intervals

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/innovate-technologies/DJ/cron"

	"github.com/innovate-technologies/DJ/at"
	"github.com/innovate-technologies/DJ/data"
	"github.com/innovate-technologies/DJ/itframe"
)

type itframeAPI interface {
	GetAllIntervals() []data.Interval
}

var api itframeAPI = itframe.New()

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
	allIntervals := api.GetAllIntervals()
	a := at.GetInstance()
	c := cron.GetInstance()

	for _, interval := range allIntervals {
		a.Add(at.Action{
			Event: "reloadQueue",
			Time:  interval.Start,
		})
		a.Add(at.Action{
			Event: "reloadQueue",
			Time:  interval.End,
		})
		for _, day := range interval.Days {
			c.Add(cron.Action{
				Event:      "reloadQueue",
				DayOfMonth: "*",
				Month:      "*",
				DayOfWeek:  strconv.Itoa(day),
				Hour:       strconv.Itoa(interval.DayStart.Hour),
				Minute:     strconv.Itoa(interval.DayStart.Minute),
			})
			c.Add(cron.Action{
				Event:      "reloadQueue",
				DayOfMonth: "*",
				Month:      "*",
				DayOfWeek:  strconv.Itoa(day),
				Hour:       strconv.Itoa(interval.DayEnd.Hour),
				Minute:     strconv.Itoa(interval.DayEnd.Minute),
			})
		}
	}
}

func getCurrentIntervals() []data.Interval {
	now := time.Now()
	day := int(now.Weekday())
	if day == 0 {
		day = 7
	}

	allIntervals := api.GetAllIntervals()
	currentIntervals := []data.Interval{}

	for _, interval := range allIntervals {
		if (interval.Forever && now.After(interval.Start)) || (now.After(interval.Start) && now.Before(interval.End)) {
			if sliceContains(interval.Days, day) {
				if inBetween(interval.DayStart.Hour, interval.DayEnd.Hour, interval.DayStart.Minute, interval.DayEnd.Minute, now) {
					currentIntervals = append(currentIntervals, interval)
				}
			}
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

func shuffle(vals []data.Song) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func sliceContains(list []int, value int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func inBetween(startHour, endHour, startMinute, endMinute int, now time.Time) bool {
	if startHour < now.Hour() && now.Hour() < endHour {
		return true
	}
	if startHour == now.Hour() && startMinute <= now.Minute() {
		return true
	}
	if endHour == now.Hour() && endMinute >= now.Minute() {
		return true
	}
	return false
}
