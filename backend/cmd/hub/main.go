package main

import (
	_ "embed"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/server"
	"github.com/getlantern/systray"
)

// go build -ldflags="-H=windowsgui" -o hub2.0.exe  ./cmd/hub
func main() {
	global.LoadEnv()
	checkRunning(8000)

	//db.Init()

	// openInBroswer()
	server.Start()

	systray.Run(setupTrayIcon, onExit)
}
