package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

// Create the authentication routes for server
func authRoutes(mux *mux.Router) {

	routes := routes{
		{
			methods{"GET"},
			"/api/verify",
			http.HandlerFunc(handlers.Verify),
			middleware.Middleware{middleware.CorsHandler},
		},
	}

	routes.add(mux)
}
