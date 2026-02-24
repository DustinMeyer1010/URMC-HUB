package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

// Creates the computer routes
func computerRoutes(mux *mux.Router) {
	routes := routes{
		{
			methods{"GET"},
			"/api/computer/{computer}/info",
			http.HandlerFunc(handlers.ComputerInfo),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
	}

	routes.add(mux)
}
