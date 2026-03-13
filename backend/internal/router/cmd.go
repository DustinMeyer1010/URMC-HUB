package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

func cmdRoutes(mux *mux.Router) {
	routes := routes{
		{
			methods{"GET"},
			"/api/ping",
			http.HandlerFunc(handlers.Ping),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/nslookup",
			http.HandlerFunc(handlers.NSLookUp),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
	}

	routes.add(mux)
}
