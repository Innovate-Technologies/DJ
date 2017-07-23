package main

import (
	"sync"
	"time"

	"fmt"

	"github.com/innovate-technologies/DJ/at"
	"github.com/innovate-technologies/DJ/cron"
	"github.com/innovate-technologies/DJ/data"
	"github.com/innovate-technologies/DJ/playlists/clocks"
	"github.com/innovate-technologies/DJ/playlists/intervals"
)

var (
	queue       = []data.Song{}
	queueMutex  = sync.Mutex{}
	currentSong data.Song
)

func init() {
	events.On("playSong", playSong)
	events.On("reloadQueue", ReloadClocks)
	events.On("queueUpdate", updateEngines)
}

// LoadClocks adds the current clock to the queue
func LoadClocks() {
	queueMutex.Lock()
	queue = append(queue, intervals.PlaceIntervals(clocks.GetSongs())...)
	queueMutex.Unlock()
	events.Emit("queueUpdate")
}

// ReloadClocks clears the queue and re-adds the songs
func ReloadClocks() {
	queueMutex.Lock()
	queue = []data.Song{}
	queueMutex.Unlock()
	LoadClocks()
	UpdateTimers()
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

func UpdateTimers() {
	cron.GetInstance().RemoveAll()
	at.GetInstance().RemoveAll()

	go clocks.SetReloads()
	go intervals.SetReloads()
}

func playSong(song data.Song) {
	queueMutex.Lock()
	if song.ID == queue[0].ID {
		currentSong = song
		queue = queue[1:]
	}
	queueMutex.Unlock()
}

func updateEngines() {
	fmt.Println("New songs, calling engine!")
	queueMutex.Lock()
	engine.PutQueue(queue)
	queueMutex.Unlock()
}
