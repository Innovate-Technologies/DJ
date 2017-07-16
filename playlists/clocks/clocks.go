package clocks

import (
	"math/rand"
	"time"

	"github.com/innovate-technologies/DJ/data"
	"github.com/innovate-technologies/DJ/itframe"
)

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

func getCurrentClock() data.Clock {
	now := time.Now()

	day := int(now.Weekday())
	if day == 0 {
		day = 7
	}

	allClocks := itframe.GetAllClocks()

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
	allSongs := itframe.GetAllSongsForTag(tag)
	shuffle(allSongs)
	if len(allSongs) >= num {
		return allSongs[:num]
	}

	for len(allSongs) < num {
		allSongs = append(allSongs, allSongs[rand.Intn(len(allSongs)-1)])
	}
	return allSongs
}

func shuffle(slc []data.Song) {
	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
}
