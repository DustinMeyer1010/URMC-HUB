package router

import (
	"github.com/gorilla/mux"
)

// Create the routes for the backend
// "/" is special as it will handle all of the react routes
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
