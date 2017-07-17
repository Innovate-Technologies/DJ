package main

import (
	"encoding/json"
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
)

func startServer() {
	io, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	io.Of("/queueEvents").On("connection", handleQueueEvents)

	e := echo.New()
	e.Any("/socket.io/*", echo.WrapHandler(io))
	e.Logger.Fatal(e.Start(":80"))
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
