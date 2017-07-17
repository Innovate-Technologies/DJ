package main

import (
	"sync"
	"time"

	"github.com/chuckpreslar/emission"
	"github.com/innovate-technologies/DJ/data"
)

var (
	queue      = []data.Song{}
	queueMutex = sync.Mutex{}
	// Events is the global event emitter
	Events *emission.Emitter
)

func init() {
	Events.On("playSong", playSong)
}

// LoadClocks adds the current clock to the queue
func LoadClocks() {
	queueMutex.Lock()
	// append here
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
