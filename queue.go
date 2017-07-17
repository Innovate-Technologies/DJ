package main

import (
	"sync"
	"time"

	"github.com/chuckpreslar/emission"
	"github.com/innovate-technologies/DJ/data"
	"github.com/innovate-technologies/DJ/playlists/clocks"
	"github.com/innovate-technologies/DJ/playlists/intervals"
)

var (
	queue      = []data.Song{}
	queueMutex = sync.Mutex{}
	// Events is the global event emitter
	Events *emission.Emitter
)

func init() {
	Events.On("playSong", playSong)
	Events.On("relodClocks", ReloadClocks)
}

// LoadClocks adds the current clock to the queue
func LoadClocks() {
	queueMutex.Lock()
	queue = append(queue, intervals.PlaceIntervals(clocks.GetSongs())...)
	queueMutex.Unlock()
	Events.Emit("queueUpdate")
}

// ReloadClocks clears the queue and re-adds the songs
func ReloadClocks() {
	queueMutex.Lock()
	queue = []data.Song{}
	queueMutex.Unlock()
	LoadClocks()
}

// WatchClocks watches the queue to add songs when needed
func WatchClocks() {
	for {
		queueMutex.Lock()
		left := len(queue)
		queueMutex.Unlock()
		if left <= 5 {
			LoadClocks()
		}
		time.Sleep(time.Second)
	}
}

func playSong(song data.Song) {
	queueMutex.Lock()
	if song.ID == queue[0].ID {
		queue = queue[1:]
	}
	queueMutex.Unlock()
}
