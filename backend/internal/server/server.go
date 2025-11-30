package server

import (
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/router"
)

var Server *http.Server

// Creates the router and start local server
func Start() {
	logger.ServerLogger.Info("Creating Routes")
	router := router.Create()
	logger.ServerLogger.Info("Routes Created")
	address := fmt.Sprintf("0.0.0.0:%d", 8000)
	Server = &http.Server{Addr: address, Handler: router}
	go Server.ListenAndServe()
	logger.ServerLogger.Info("Server Started on http://localhost:8000")

}
