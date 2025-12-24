package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

// Creates the computer routes
func imageRoutes(mux *mux.Router) {
	routes := routes{
		{
			methods{"GET"},
			"/api/static/image/{imageName}",
			http.HandlerFunc(handlers.StaticImage),
			middleware.Middleware{middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/db/image/{imageName}",
			http.HandlerFunc(handlers.DBImage),
			middleware.Middleware{middleware.CorsHandler},
		},
	}

	routes.add(mux)
}
