package router

import (
	"github.com/gorilla/mux"
)

// Creates All routes for the server
func Create() *mux.Router {

	mux := mux.NewRouter()

	createRoutes(mux,
		reactRoutes,
		searchRoutes,
		userRoutes,
		authRoutes,
		bookmarksRoutes,
	)

	return mux

}

func createRoutes(mux *mux.Router, routes ...func(*mux.Router)) {
	for _, route := range routes {
		route(mux)
	}
}
