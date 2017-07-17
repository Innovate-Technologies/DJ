package main

import (
	"encoding/json"
	"log"

	"net/http"

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
	e.GET("/api/:key/songs/queue", getQueue)
	e.GET("/api/:key/songs/current", getCurrentSong)
	e.POST("/api/:key/songs/skip", postSkip)
	e.POST("/api/:key/clocks/reload", postReload)
	e.Any("/socket.io/*", echo.WrapHandler(io))
	e.Use(checkKey)
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

func checkKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Path()[:5] == "/api/" && c.Param("key") != conf.APIKey {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "bad API key"})
		}
		return next(c)
	}
}

func getQueue(c echo.Context) error {
	return c.JSON(http.StatusOK, queue)
}

func getCurrentSong(c echo.Context) error {
	return c.JSON(http.StatusOK, currentSong)
}

func postSkip(c echo.Context) error {
	events.Emit("skipSong")
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func postReload(c echo.Context) error {
	events.Emit("reloadQueue")
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
