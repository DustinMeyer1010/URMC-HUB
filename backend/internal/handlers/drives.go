package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
)

func DriveAccess(w http.ResponseWriter, r *http.Request) {

	var groups []string

	if err := json.NewDecoder(r.Body).Decode(&groups); err != nil {
		http.Error(w, "INVALID_BODY", http.StatusBadRequest)
		return
	}

	drives, statusError := service.GetDriveAccess(groups)

	if statusError != nil {
		http.Error(w, statusError.ErrorType, statusError.HttpStatus)
		w.Write([]byte(statusError.Error))
		return
	}

	data, err := json.Marshal(drives)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
