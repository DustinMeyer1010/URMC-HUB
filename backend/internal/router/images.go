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
			"/api/image/{imageName}",
			http.HandlerFunc(handlers.StaticPhotos),
			middleware.Middleware{middleware.CorsHandler},
		},
	}

	routes.add(mux)
}
