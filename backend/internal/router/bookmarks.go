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
			"/api/bookmark/{username}",
			http.HandlerFunc(handlers.AddBookmark),
			middleware.Middleware{
				middleware.IsAuthorized,
				middleware.IsAuthorizedUser,
				middleware.CorsHandler,
			},
		},
		{
			methods{"GET"},
			"/api/bookmarks/{username}",
			http.HandlerFunc(handlers.GetBookForAgent),
			middleware.Middleware{
				middleware.IsAuthorized,
				middleware.IsAuthorizedUser,
				middleware.CorsHandler,
			},
		},
		{
			methods{"GET"},
			"/api/generic/bookmarks",
			http.HandlerFunc(handlers.GetGeneralBookmarks),
			middleware.Middleware{
				middleware.IsAuthorized,
				middleware.CorsHandler,
			},
		},
		{
			methods{"GET"},
			"/api/bookmarks/all/agents",
			http.HandlerFunc(handlers.GetAgentsWithBookmarks),
			middleware.Middleware{
				middleware.IsAuthorized,
				middleware.CorsHandler,
			},
		},
	}

	routes.add(mux)

}
