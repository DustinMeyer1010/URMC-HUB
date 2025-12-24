package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

// Creates bookmarks routes
func utilsRoutes(mux *mux.Router) {

	routes := routes{
		{
			methods{"POST", "OPTIONS"},
			"/api/directory",
			http.HandlerFunc(handlers.OpenDirectory),
			middleware.Middleware{
				middleware.IsAuthorized,
				middleware.CorsHandler,
			},
		},
	}

	routes.add(mux)

}
