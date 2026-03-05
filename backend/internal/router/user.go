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
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
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
			"/api/user",
			http.HandlerFunc(handlers.GroupAddToUser),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			// * NEW
			methods{"DELETE", "OPTIONS"},
			"/api/user",
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
		{
			// Deprecated: To be Replace with GET /api/user
			methods{"GET"},
			"/api/user/{username}",
			http.HandlerFunc(handlers.PullUserInformation),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			// Deprecated: To be Replace with GET /api/user/groups
			methods{"GET"},
			"/api/user/{username}/memberof",
			http.HandlerFunc(handlers.GetMemberOf),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},

		{
			// Deprecated: To be Replace with GET /api/user/lockout
			methods{"GET"},
			"/api/user/{username}/lockout",
			http.HandlerFunc(handlers.LockOutStatus),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},

		{
			// Deprecated: To be Replace with GET /api/user/drives
			methods{"POST", "OPTIONS"},
			"/api/drive/access",
			http.HandlerFunc(handlers.DriveAccess),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},

		{
			// Deprecated: To be Replace with POST /api/user
			methods{"POST", "OPTIONS"},
			"/api/user/{username}/memberof",
			http.HandlerFunc(handlers.AddGroup),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			// Deprecated: To be Replace with DELETE /api/user
			methods{"DELETE", "OPTIONS"},
			"/api/user/{username}/memberof",
			http.HandlerFunc(handlers.RemoveGroup),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
	}

	routes.add(mux)
}
