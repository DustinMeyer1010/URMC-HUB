package middleware

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
)

func IsAuthorized(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if global.Username == "" || global.Password == "" {
				http.Error(w, "Unauthorized (Not Logged In)", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
}
