package router

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/gorilla/mux"
)

// Create the authentication routes for server
func authRoutes(mux *mux.Router) {
	mux.HandleFunc("/verify", handlers.Verify).Methods("GET")
}
