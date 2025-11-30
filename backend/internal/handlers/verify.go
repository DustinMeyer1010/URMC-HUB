package handlers

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
)

// Verify that the users creds are still valid
func Verify(w http.ResponseWriter, r *http.Request) {

	cError := ad.Verify()

	if cError != nil {
		logger.LogRequestInfo(r.Method, r.URL.Path)
		logger.Error(cError)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(cError.Msg))
		return
	}
	w.WriteHeader(http.StatusOK)
}
