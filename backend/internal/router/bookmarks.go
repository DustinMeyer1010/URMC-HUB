package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/gorilla/mux"
)

func bookmarksRoutes(mux *mux.Router) {
	mux.HandleFunc("/bookmarks", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	mux.HandleFunc("/bookmarks", handlers.AddBookmark).Methods("POST")
	mux.HandleFunc("/bookmarks", func(w http.ResponseWriter, r *http.Request) {}).Methods("UPDATE")
	mux.HandleFunc("/bookmarks/{id}", handlers.RemoveBookmark).Methods("DELETE")

}
