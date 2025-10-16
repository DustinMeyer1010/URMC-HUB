package router

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/gorilla/mux"
)

func searchRoutes(mux *mux.Router) {
	mux.HandleFunc("/search/users/{searchValue}", handlers.UserSearch).Methods("GET")
	mux.HandleFunc("/search/groups/{searchValue}", handlers.GroupSearch).Methods("GET")
	mux.HandleFunc("/search/computers/{searchValue}", handlers.ComputerSearch).Methods("GET")
	mux.HandleFunc("/search/printers/{searchValue}", handlers.PrinterSearch).Methods("GET")
	mux.HandleFunc("/search/sharedrive/{searchValue}", handlers.ShareDriveShearch).Methods("GET")
	mux.HandleFunc("/search/all/{searchValue}", handlers.AllSearch).Methods("GET")
}
