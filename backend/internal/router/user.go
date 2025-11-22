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
			"/api/user/info/{username}",
			http.HandlerFunc(handlers.PullUserInformation),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/lockout/{username}",
			http.HandlerFunc(handlers.LockOutStatus),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"POST", "OPTIONS"},
			"/api/user/login",
			http.HandlerFunc(handlers.Login),
			middleware.Middleware{middleware.CorsHandler},
		},
		{
			methods{"POST", "OPTIONS"},
			"/api/drive/access",
			http.HandlerFunc(handlers.DriveAccess),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"POST", "OPTIONS"},
			"/api/user/group/add",
			http.HandlerFunc(handlers.AddGroup),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"POST", "OPTIONS"},
			"/api/user/group/remove",
			http.HandlerFunc(handlers.RemoveGroup),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"POST", "OPTION"},
			"/api/users/bulk/lookup/file",
			http.HandlerFunc(handlers.BulkUserSearchFile),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"POST", "OPTION"},
			"/api/users/bulk/lookup",
			http.HandlerFunc(handlers.BulkUserSearch),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
	}

	routes.add(mux)
}
