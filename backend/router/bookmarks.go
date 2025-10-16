package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func bookmarksRoutes(mux *mux.Router) {
	mux.HandleFunc("/bookmarks", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	mux.HandleFunc("/bookmarks", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")
	mux.HandleFunc("/bookmarks", func(w http.ResponseWriter, r *http.Request) {}).Methods("UPDATE")
	mux.HandleFunc("/bookmarks", func(w http.ResponseWriter, r *http.Request) {}).Methods("DELETE")

}
