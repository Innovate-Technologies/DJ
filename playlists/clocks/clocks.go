package clocks

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/innovate-technologies/DJ/cron"
	"github.com/innovate-technologies/DJ/data"
	"github.com/innovate-technologies/DJ/itframe"
)

type itframeAPI interface {
	GetAllClocks() []data.Clock
	GetAllSongsForTag(tag string) []data.Song
}

var api itframeAPI = itframe.New()

// GetSongs returns 100 songs for the current clock
func GetSongs() (songs []data.Song) {
	songs = []data.Song{}
	clock := getCurrentClock()

	for _, tag := range clock.Tags {
		songs = append(songs, getSongsForTag(tag.Tag, tag.Percent)...)
	}
	shuffle(songs)
	return
}

// SetReloads sets the reload events when a new clock has to start
func SetReloads() {
	allClocks := api.GetAllClocks()
	c := cron.GetInstance()

	for _, clock := range allClocks {
		c.Add(cron.Action{
			DayOfMonth: "*",
			Month:      "*",
			DayOfWeek:  strconv.FormatInt(int64(clock.Start.DayOfWeek), 10),
			Hour:       strconv.FormatInt(int64(clock.Start.Hour), 10),
			Minute:     strconv.FormatInt(int64(clock.Start.Minute), 10),
			Event:      "reloadQueue",
		})
	}
}

func getCurrentClock() data.Clock {
	now := time.Now()

	day := int(now.Weekday())
	if day == 0 {
		day = 7
	}

	allClocks := api.GetAllClocks()

	for _, clock := range allClocks {
		if clock.Start.DayOfWeek < day && clock.End.DayOfWeek > day {
			// ] day [
			return clock
		}
		if clock.Start.DayOfWeek == day || clock.End.DayOfWeek == day {
			// [ day ]
			if (clock.Start.DayOfWeek == day && clock.Start.Hour <= now.Hour()) || (clock.End.DayOfWeek == day && clock.End.Hour >= now.Hour()) {
				// check end minutes
				// [ day ] [ hour ]
				if clock.Start.Hour < now.Hour() && clock.End.Hour > now.Hour() {
					// ] hour [
					return clock
				} else if (clock.Start.Hour == now.Hour() && clock.Start.Minute >= now.Minute()) || (clock.End.Hour == now.Hour() && clock.End.Minute >= now.Minute()) {
					// [ day ] [ hour ] [ minute ]
					return clock
				}
			}
		}
	}

	return data.Clock{}
}

func getSongsForTag(tag string, num int) []data.Song {
	allSongs := api.GetAllSongsForTag(tag)
	shuffle(allSongs)
	if len(allSongs) >= num {
		return allSongs[:num]
	}

	for len(allSongs) < num {
		if len(allSongs) > 1 {
			allSongs = append(allSongs, allSongs[rand.Intn(len(allSongs)-1)])
		} else {
			allSongs = append(allSongs, allSongs[0])
		}
	}
	return allSongs
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
