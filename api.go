package main

import (
	"encoding/json"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func startServer() {
	io, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	io.Of("/queueEvents").On("connection", handleQueueEvents)

	http.Handle("/socket.io/", io)

	log.Fatal(http.ListenAndServe(":80", nil))
}

func handleQueueEvents(socket socketio.Socket) {
	socket.On("key", func(key string) {
		if key == conf.APIKey {
			go sendCurrentSong(socket)
			go sendQueue(socket)

			events.On("queueUpdate", func() {
				go sendCurrentSong(socket)
				go sendQueue(socket)
			})

		} else {
			socket.Emit("error", "invalid key")
		}
	})
}

func sendCurrentSong(socket socketio.Socket) {
	queueMutex.Lock()
	jsonOut, _ := json.Marshal(currentSong)
	socket.Emit("currentSong", string(jsonOut))
	queueMutex.Unlock()
}

func sendQueue(socket socketio.Socket) {
	queueMutex.Lock()
	jsonOut, _ := json.Marshal(queue)
	socket.Emit("currentSong", string(jsonOut))
	queueMutex.Unlock()
}
