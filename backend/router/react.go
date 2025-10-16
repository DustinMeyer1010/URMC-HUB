package router

import "github.com/gorilla/mux"

func reactRoutes(mux *mux.Router) {
	mux.HandleFunc("/", reactHandler).Methods("GET")
}
