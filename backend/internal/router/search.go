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
			middleware.Middleware{middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/search/groups/{searchValue}",
			http.HandlerFunc(handlers.GroupSearch),
			middleware.Middleware{middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/search/computer/{searchValue}",
			http.HandlerFunc(handlers.ComputerSearch),
			middleware.Middleware{middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/search/printer/{searchValue}",
			http.HandlerFunc(handlers.PrinterSearch),
			middleware.Middleware{middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/search/sharedrive/{searchValue}",
			http.HandlerFunc(handlers.ShareDriveSearch),
			middleware.Middleware{middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/search/all/{searchValue}",
			http.HandlerFunc(handlers.AllSearch),
			middleware.Middleware{middleware.CorsHandler},
		},
	}

	routes.add(mux)

}
