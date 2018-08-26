package main

import (
	"log"
	"strconv"
	"strings"

	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
)

func startServer() {
	io, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	io.On("connection", handleQueueEvents)
	e := echo.New()
	e.GET("/", getRoot)
	e.GET("/api/:key/songs/queue", getQueue)
	e.GET("/api/:key/songs/current", getCurrentSong)
	e.POST("/api/:key/songs/skip", postSkip)
	e.POST("/api/:key/clocks/reload", postReload)
	e.Any("/socket.io/*", echo.WrapHandler(io))
	e.Use(checkKey)
	e.Use(socketioCORS)
	e.Debug = false
	e.HideBanner = true
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

			events.On("timeRemaining", func(time int) {
				go socket.Emit("timeRemaining", strconv.FormatInt(int64(time), 10))
			})

		} else {
			socket.Emit("error", "invalid key")
		}
	})
}

func sendCurrentSong(socket socketio.Socket) {
	queueMutex.Lock()
	socket.Emit("currentSong", currentSong)
	queueMutex.Unlock()
}

func sendQueue(socket socketio.Socket) {
	queueMutex.Lock()
	socket.Emit("queue", queue)
	queueMutex.Unlock()
}

func socketioCORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasPrefix(c.Path(), "/socket.io/") {
			if origin := c.Request().Header.Get("Origin"); origin != "" {
				c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
				c.Response().Header().Set("Access-Control-Allow-Origin", origin)
			}
		}
		return next(c)
	}
}

func checkKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasPrefix(c.Path(), "/api/") && c.Param("key") != conf.APIKey {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "bad API key"})
		}
		return next(c)
	}
}

func getRoot(c echo.Context) error {
	return c.String(http.StatusOK, "DJ Server")
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
