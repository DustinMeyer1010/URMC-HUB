package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/server"
	"github.com/getlantern/systray"
)

//go:embed URMC.ico
var APPICON []byte

func main() {
	global.LoadEnv()
	//db.Init()

	if checkRunning(8000) {
		fmt.Println("Server already running")
		os.Exit(1)
	}

	server.Start()
	systray.Run(setupTrayIcon, onExit)
}
