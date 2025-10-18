package handlers

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
)

// Verify that the users creds are still valid
func Verify(w http.ResponseWriter, r *http.Request) {

	err := ad.Verify()

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
