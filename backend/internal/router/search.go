package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

// Creates all the search routes
func searchRoutes(mux *mux.Router) {

	routes := routes{
		{
			methods{"GET"},
			"/api/search/users/{searchValue}",
			http.HandlerFunc(handlers.UserSearch),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/search",
			http.HandlerFunc(handlers.SearchAll),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/search/groups/{searchValue}",
			http.HandlerFunc(handlers.GroupSearch),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/search/computers/{searchValue}",
			http.HandlerFunc(handlers.ComputerSearch),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/search/printers/{searchValue}",
			http.HandlerFunc(handlers.PrinterSearch),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/search/drives/{searchValue}",
			http.HandlerFunc(handlers.ShareDriveSearch),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/search/all/{searchValue}",
			http.HandlerFunc(handlers.AllSearch),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
	}

	routes.add(mux)

}
