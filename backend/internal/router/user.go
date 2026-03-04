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
			"/api/user",
			http.HandlerFunc(handlers.GetUser),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/user/attributes",
			http.HandlerFunc(handlers.GetUserAvaiableAttributes),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/user/drives",
			http.HandlerFunc(handlers.GetUserAvaiableAttributes),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/user/groups",
			http.HandlerFunc(handlers.GetUserAvaiableAttributes),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/user/lockout",
			http.HandlerFunc(handlers.GetUserAvaiableAttributes),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/user/{username}",
			http.HandlerFunc(handlers.PullUserInformation),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/user/{username}/memberof",
			http.HandlerFunc(handlers.GetMemberOf),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},

		{
			methods{"GET"},
			"/api/user/{username}/lockout",
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
			"/api/user/{username}/memberof",
			http.HandlerFunc(handlers.AddGroup),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"DELETE", "OPTIONS"},
			"/api/user/{username}/memberof",
			http.HandlerFunc(handlers.RemoveGroup),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
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
	}

	routes.add(mux)
}
