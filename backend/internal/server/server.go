package server

import (
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/router"
	"github.com/LostProgrammer1010/URMC-HUB/internal/utils"
	"github.com/go-toast/toast"
)

var Server *http.Server

// Creates the router and start local server
func Start() {
	router := router.Create()
	address := fmt.Sprintf("0.0.0.0:%d", 8000)
	Server = &http.Server{Addr: address, Handler: router}
	go serveServer()
	notification := toast.Notification{
		AppID:    "URMC-HUB",
		Title:    "Notification",
		Icon:     utils.GetIconPath(),
		Message:  "The application has started",
		Duration: toast.Long,
	}

	notification.Push()
	fmt.Println("Server Running")
	logger.Info("Server Started Successfull on http://localhost:8000")

}

func serveServer() {
	err := Server.ListenAndServe()

	if err != nil {
		notification := toast.Notification{
			AppID:    "URMC-HUB",
			Title:    "Warning",
			Message:  fmt.Sprintf("Failed to start server.\nError:\n%s", err.Error()),
			Duration: toast.Short,
		}

		notification.Push()

		panic("Server failed to start")
	}

}
