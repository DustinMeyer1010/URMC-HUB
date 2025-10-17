package router

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/gorilla/mux"
)

func authRoutes(mux *mux.Router) {
	mux.HandleFunc("/verify", handlers.Verify).Methods("GET")
}
