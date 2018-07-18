package main

import (
	"fmt"
	"os"

	"github.com/chuckpreslar/emission"
	"github.com/innovate-technologies/DJ/at"
	"github.com/innovate-technologies/DJ/config"
	"github.com/innovate-technologies/DJ/cron"
	dummyengine "github.com/innovate-technologies/dummy-dj-engine"
)

var (
	username = os.Getenv("username")
	conf     config.Config
	events   = emission.NewEmitter()
	engine   Engine
)

func main() {
	fmt.Print("     _____        ___    \n    /  /::\\      /  /\\   \n   /  /:/\\:\\    /  /:/   \n  /  /:/  \\:\\  /__/::\\   \n /__/:/ \\__\\:| \\__\\/\\:\\  \n \\  \\:\\ /  /:/    \\  \\:\\ \n  \\  \\:\\  /:/      \\__\\:\\\n   \\  \\:\\/:/       /  /:/\n    \\  \\::/       /__/:/ \n     \\__\\/        \\__\\/  \n                         \n")
	fmt.Println("Copyright 2018 SHOUTca.st")
	fmt.Println("=========================")

	if username == "" {
		fmt.Println("No username provided")
		return
	}

	conf = *config.GetConfig()

	initQueueEvents()

	engine = dummyengine.New()

	// init At and cron
	at.Events = events
	cron.Events = events
	at.GetInstance()
	cron.GetInstance()

	ReloadClocks()

	go engine.Start(events)

	startServer()
}
