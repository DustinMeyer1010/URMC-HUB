package main

import (
	"fmt"
	"net"

	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/server"
	"github.com/getlantern/systray"
)

func onExit() {
	logger.Info("Server Shut Down")
}

func setupTrayIcon() {

	systray.SetIcon(APPICON)
	systray.SetTitle("URMC-HUB Server")
	systray.SetTooltip("URMC-HUB Server")
	quitMenuItem := systray.AddMenuItem("Quit", "Exit the application")

	go func() {
		<-quitMenuItem.ClickedCh
		systray.Quit()
		server.Server.Close()

	}()
}

func checkRunning(port int) bool {
	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return true
	}
	defer listener.Close()
	return false
}
