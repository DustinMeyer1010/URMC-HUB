package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

// Creates bookmarks routes
func bookmarksRoutes(mux *mux.Router) {

	routes := routes{
		{
			methods{"GET", "POST", "UPDATE"},
			"/bookmarks",
			http.HandlerFunc(handlers.Bookmarks),
			middleware.Middleware{middleware.CorsHandler},
		},
		{
			methods{"DELETE"},
			"bookmarks/{id}",
			http.HandlerFunc(handlers.Bookmarks),
			middleware.Middleware{middleware.CorsHandler},
		},
	}

	routes.add(mux)

}
