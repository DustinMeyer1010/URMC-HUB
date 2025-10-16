package router

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/gorilla/mux"
)

func userRoutes(mux *mux.Router) {
	mux.HandleFunc("/lockout/{username}", handlers.LockOutStatus).Methods("GET")
	mux.HandleFunc("/user/{URID}", handlers.PullUserInformation).Methods("GET")
	mux.HandleFunc("/user/login", handlers.Login).Methods("POST")
}
