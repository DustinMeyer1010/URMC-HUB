package main

import (
	_ "embed"
	"os"

	"github.com/LostProgrammer1010/URMC-HUB/internal/db"
	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/server"
	"github.com/getlantern/systray"
)

//go:embed URMC.ico
var APPICON []byte

func main() {
	global.LoadEnv()

	db.Init()

	if checkRunning(8000) {
		logger.Error("Server already running on 8000")
		os.Exit(1)
	}

	server.Start()
	systray.Run(setupTrayIcon, onExit)
}
