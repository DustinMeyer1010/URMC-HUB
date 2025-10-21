package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Creates All routes for the server
func Create() *mux.Router {

	mux := mux.NewRouter()

	createRoutes(mux,
		searchRoutes,
		userRoutes,
		authRoutes,
		bookmarksRoutes,
		reactRoutes,
	)

	mux.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("404 Page"))
	})

	return mux

}

// Helper function to seperate routes into seperate files
func createRoutes(mux *mux.Router, routes ...func(*mux.Router)) {
	for _, route := range routes {
		route(mux)
	}
}
