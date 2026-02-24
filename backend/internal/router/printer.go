package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

// Creates the printer routes
func printerRoutes(mux *mux.Router) {

	routes := routes{
		{
			methods{"GET"},
			"/api/printer/info/{server}",
			http.HandlerFunc(handlers.PrinterInformation),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/printer/ping/{ip}",
			http.HandlerFunc(handlers.PingPrinter),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			methods{"GET"},
			"/api/printer/related/{ip}",
			http.HandlerFunc(handlers.RelatedPrinters),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
	}

	routes.add(mux)

}
