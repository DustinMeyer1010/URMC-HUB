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
			methods{"POST", "OPTION"},
			"/api/add/bookmark",
			http.HandlerFunc(handlers.AddBookmark),
			middleware.Middleware{middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/generate/generic/bookmarks",
			http.HandlerFunc(handlers.GenerateGenericBookmarks),
			middleware.Middleware{
				middleware.CorsHandler,
				middleware.AdminCheck,
			},
		},
	}

	routes.add(mux)

}
