package main

import (
	"fmt"
	"os"

	"github.com/chuckpreslar/emission"
	"github.com/innovate-technologies/DJ/config"
	"github.com/innovate-technologies/DJ/itframe"
)

var (
	username = os.Getenv("username")
	conf     config.Config
	events   = emission.NewEmitter()
)

func main() {
	fmt.Print("     _____        ___    \n    /  /::\\      /  /\\   \n   /  /:/\\:\\    /  /:/   \n  /  /:/  \\:\\  /__/::\\   \n /__/:/ \\__\\:| \\__\\/\\:\\  \n \\  \\:\\ /  /:/    \\  \\:\\ \n  \\  \\:\\  /:/      \\__\\:\\\n   \\  \\:\\/:/       /  /:/\n    \\  \\::/       /__/:/ \n     \\__\\/        \\__\\/  \n                         \n")
	fmt.Println("Copyright 2017 SHOUTca.st")
	fmt.Println("=========================")

	if username == "" {
		fmt.Println("No username provided")
		return
	}

	conf = itframe.GetConfig(username)

}
