package middleware

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
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
