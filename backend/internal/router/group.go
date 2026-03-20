package router

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/handlers"
	"github.com/LostProgrammer1010/URMC-HUB/internal/middleware"
	"github.com/gorilla/mux"
)

// Creates the group routes
func groupRoutes(mux *mux.Router) {
	routes := routes{
		{
			// * NEW
			methods{"GET"},
			"/api/group",
			http.HandlerFunc(handlers.GetGroup),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			// * NEW
			methods{"GET"},
			"/api/group/attributes",
			http.HandlerFunc(handlers.GetGroupAvaiableAttributes),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			// * NEW
			methods{"GET"},
			"/api/group/members",
			http.HandlerFunc(handlers.GetGroupMembers),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			// * NEW
			methods{"POST", "OPTION"},
			"/api/group/members",
			http.HandlerFunc(handlers.UsersAddTOGroup),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		{
			// * NEW
			methods{"DELETE", "OPTION"},
			"/api/group/members",
			http.HandlerFunc(handlers.UsersRemoveFromGroup),
			middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
		},
		/*
			{
				methods{"GET", "OPTION"},
				"/api/group/{group}/members/excel",
				http.HandlerFunc(handlers.GetAllMembersExcel),
				middleware.Middleware{middleware.IsAuthorized, middleware.CorsHandler},
			},
		*/
	}

	routes.add(mux)
}
