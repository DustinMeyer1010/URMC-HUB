package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

// Create all the user routes
func userRoutes(mux *mux.Router) {

	routes := routes{
		{
			methods{"GET"},
			"/user/info/{username}",
			http.HandlerFunc(handlers.PullUserInformation),
			middleware.Middleware{middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/lockout/{username}",
			http.HandlerFunc(handlers.LockOutStatus),
			middleware.Middleware{middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/user/login",
			http.HandlerFunc(handlers.Login),
			middleware.Middleware{middleware.CorsHandler},
		},
	}

	routes.add(mux)
}
