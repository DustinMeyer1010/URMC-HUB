package middleware

import (
	"net/http"
)

func CheckForDNQuery(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			dn := r.URL.Query().Get("dn")

			if dn == "" {
				http.Error(w, "User DistingusishedName must be included in this endpoint. This can be added as a query parameter of dn", http.StatusBadRequest)
				return
			}
			next.ServeHTTP(w, r)
		})
}
