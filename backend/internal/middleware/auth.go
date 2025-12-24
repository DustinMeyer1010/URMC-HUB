package middleware

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/gorilla/mux"
)

// Checks to make sure the person is logged in before doing request
func IsAuthorized(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if global.Username == "" || global.Password == "" {
				http.Error(w, "UNAUTHORIZED LOGIN STATUS FALSE", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
}

func IsAuthorizedUser(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			username := vars["username"]
			if global.Username != username {
				http.Error(w, "UNAUTHORIZED_NOT_LOGGED_IN_USER", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
}
