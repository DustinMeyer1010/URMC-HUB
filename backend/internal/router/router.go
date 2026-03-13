package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Creates All routes for the server
func Create() *mux.Router {

	mux := mux.NewRouter()
	mux.UseEncodedPath()

	// Add a function to list with the routes that need to be added
	createRoutes(
		mux,
		searchRoutes,
		printerRoutes,
		userRoutes,
		authRoutes,
		computerRoutes,
		groupRoutes,
		cmdRoutes,
		bookmarksRoutes,
		imageRoutes,
		driveRoutes,
		utilsRoutes,
		reactSvelte,
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
