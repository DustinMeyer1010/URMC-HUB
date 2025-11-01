package middleware

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
)

func AdminCheck(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if global.Username != global.ADMIN {
				http.Error(w, "Admin endpoint only", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
}
