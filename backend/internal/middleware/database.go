package middleware

import (
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/db"
	"github.com/LostProgrammer1010/URMC-HUB/internal/utils"
	"github.com/gorilla/mux"
)

func CheckForDatabase(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			username := vars["username"]
			if !utils.FileExist(fmt.Sprintf("%s/%s.db", db.DBLocation, username)) {
				http.Error(w, "NO_DATABASE_FOR_USER", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
}
