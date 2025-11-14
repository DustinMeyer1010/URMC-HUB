package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

func computerRoutes(mux *mux.Router) {
	routes := routes{
		{
			methods{"GET"},
			"/api/computer/info/{computer}",
			http.HandlerFunc(handlers.ComputerInfo),
			middleware.Middleware{middleware.CorsHandler},
		},
	}

	routes.add(mux)
}
