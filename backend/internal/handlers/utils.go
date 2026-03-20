package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
)

func OpenDirectory(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var directory models.Directory

	if jsonError := decoder.Decode(&directory); jsonError != nil {
		http.Error(w, "INVALID_BODY", http.StatusBadRequest)
		return
	}

	cError := service.OpenDirectory(directory)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	w.WriteHeader(http.StatusOK)
}
