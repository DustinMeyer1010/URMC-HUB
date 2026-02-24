package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

// Creates the group routes
func driveRoutes(mux *mux.Router) {
	routes := routes{
		{
			methods{"GET"},
			"/api/group/{group}/members",
			http.HandlerFunc(handlers.DriveInfo),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
	}

	routes.add(mux)
}
