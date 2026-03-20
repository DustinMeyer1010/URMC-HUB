package handlers

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
)

// Verify that the users creds are still valid
func Verify(w http.ResponseWriter, r *http.Request) {

	cError := ad.Verify()

	if e := errs.IsApiError(cError); e != nil {
		logger.LogRequestInfo(r.Method, r.URL.Path)
		logger.Error(e)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(e.Msg))
		return
	}
	w.WriteHeader(http.StatusOK)
}
