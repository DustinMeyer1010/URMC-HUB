package middleware

import (
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
)

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

func ValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("username")

			if err != nil {
				if err == http.ErrNoCookie {
					http.Error(w, "USERNAME COOKIE NOT FOUND", http.StatusUnauthorized)
					return
				}
				http.Error(w, "ERROR READING COOKIES", http.StatusBadRequest)
				return
			}

			fmt.Println(cookie.Value, global.Username)

			if global.Username != cookie.Value {
				http.Error(w, "UNAUTHORIZED CAN'T MAKE CHANGE TO ANOTHER USERS ACCOUNT", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
}
