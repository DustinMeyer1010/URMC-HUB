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
	router := router.Create()
	address := fmt.Sprintf("0.0.0.0:%d", 8000)
	Server = &http.Server{Addr: address, Handler: router}
	go Server.ListenAndServe()
	fmt.Println("Server Running")
	logger.Info("Server Started Successfull on http://localhost:8000")

}
