package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

type methods []string

type routes []route

// Adds all the routes to the mux Router that is provided
func (r routes) add(mux *mux.Router) {
	for _, route := range r {
		mux.Handle(
			route.path,
			route.middleware.MiddlewareChain(
				route.mainHandler,
			),
		).Methods(route.methods...)
	}
}

type route struct {
	methods     []string
	path        string
	mainHandler http.Handler
	middleware  middleware.Middleware
}
