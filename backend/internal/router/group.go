package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

func groupRoutes(mux *mux.Router) {
	routes := routes{
		{
			methods{"POST", "OPTION"},
			"/api/{group}/members",
			http.HandlerFunc(handlers.AddUsersToGroup),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"DELETE", "OPTION"},
			"/api/{group}/members",
			http.HandlerFunc(handlers.RemoveUsersFromGroup),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET", "OPTION"},
			"/api/{group}/info",
			http.HandlerFunc(handlers.PullUserInformation),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		/*
			{
				methods{"GET", "OPTION"},
				"/api/{group}/members",
				http.HandlerFunc(handlers.RemoveUsersFromGroup),
				middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
			},
		*/
	}

	routes.add(mux)
}
