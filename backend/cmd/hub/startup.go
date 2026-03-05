package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/server"
	"github.com/getlantern/systray"
)

func onExit() {
	logger.Info("Server Shut Down")
}

// Put the application in the hidden icon tray. Giving the users
// the ability to quit the application when something is broken.
func setupTrayIcon() {

	systray.SetIcon(APPICON)
	systray.SetTitle("URMC-HUB Server")
	systray.SetTooltip("URMC-HUB Server")
	openApp := systray.AddMenuItem("Open App", "Opens the application at localhost:8000")
	openLogs := systray.AddMenuItem("Logs", "URMC-HUB Logs")
	loginStatus := systray.AddMenuItem("Login Status: False", "")
	loginStatus.Disable()
	quitMenuItem := systray.AddMenuItem("Quit", "Exit the application")

	// Handles the exiting of the program
	go func() {
		<-quitMenuItem.ClickedCh
		systray.Quit()
		server.Server.Close()

	}()

	// Open the application in the browser if needed
	go func() {
		for range openApp.ClickedCh {
			openInBroswer()
		}
	}()

	// Opens the log files of the application
	go func() {
		for range openLogs.ClickedCh {
			openLogsInFiles()
		}
	}()

	// View of the login status of the application
	go func() {
		update := false
		for {
			if !update {
				if global.Username != "" || global.Password != "" {
					loginStatus.SetTitle("Login Status: True")
					update = true
				}
			}

		}
	}()

}

// Verifies that the server is not already running.
// If the server is already running then it close the current
// instances
func checkRunning(port int) bool {
	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return true
	}
	defer listener.Close()
	return false
}

// Opens the application in the broswer automatically rather than
// manaully going to the application page
func openInBroswer() {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", "http://localhost:8000/search"}
	}
	fmt.Println(exec.Command(cmd, args...).Start())
}

func openLogsInFiles() {
	var cmd string
	var args []string

	fmt.Println(filepath.Join(os.Getenv("LOCALAPPDATA"), "URMC-HUB"))

	switch runtime.GOOS {
	case "windows":
		cmd = "explorer"
		args = []string{filepath.Join(os.Getenv("LOCALAPPDATA"), "URMC-HUB")}
	}
	fmt.Println(exec.Command(cmd, args...).Start())
}
