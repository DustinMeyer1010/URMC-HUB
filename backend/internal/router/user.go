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
			// * NEW
			methods{"GET"},
			"/api/user",
			http.HandlerFunc(handlers.GetUser),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler, middleware.CheckForDNQuery},
		},
		{
			// * NEW
			methods{"GET"},
			"/api/user/attributes",
			http.HandlerFunc(handlers.GetUserAvaiableAttributes),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			// * NEW
			methods{"GET"},
			"/api/user/drives",
			http.HandlerFunc(handlers.GetUserDrives),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			// * NEW
			methods{"GET"},
			"/api/user/groups",
			http.HandlerFunc(handlers.GetUserGroups),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			// * NEW
			methods{"GET"},
			"/api/user/lockout",
			http.HandlerFunc(handlers.GetUserLockoutStatus),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			// * NEW
			methods{"POST", "OPTIONS"},
			"/api/user/group",
			http.HandlerFunc(handlers.GroupAddToUser),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			// * NEW
			methods{"DELETE", "OPTIONS"},
			"/api/user/group",
			http.HandlerFunc(handlers.GroupRemoveFromUser),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			// * NEW
			methods{"GET"},
			"/api/user/lockout",
			http.HandlerFunc(handlers.GetUserLockoutStatus),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"POST", "OPTIONS"},
			"/api/user/login",
			http.HandlerFunc(handlers.Login),
			middleware.Middleware{middleware.CorsHandler},
		},
		/*
			{
				methods{"POST", "OPTION"},
				"/api/users/bulk-lookup-file",
				http.HandlerFunc(handlers.BulkUserSearchFile),
				middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
			},
			{
				methods{"POST", "OPTION"},
				"/api/users/bulk-lookup",
				http.HandlerFunc(handlers.BulkUserSearch),
				middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
			},
		*/
	}

	routes.add(mux)
}
