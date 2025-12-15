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
			middleware.Middleware{
				middleware.IsAuthorized,
				middleware.ValidateUser,
				middleware.CorsHandler,
			},
		},
		{
			methods{"GET"},
			"/api/generic/bookmarks",
			http.HandlerFunc(handlers.GetGeneralBookmarks),
			middleware.Middleware{
				middleware.CorsHandler,
				middleware.AdminCheck,
			},
		},
		{
			methods{"GET"},
			"/api/bookmarks/all/agents",
			http.HandlerFunc(handlers.GetAgentsWithBookmarks),
			middleware.Middleware{
				middleware.CorsHandler,
				middleware.AdminCheck,
			},
		},
	}

	routes.add(mux)

}
