package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
)

func DriveAccess(w http.ResponseWriter, r *http.Request) {

	var groups []string

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&groups); err != nil {
		http.Error(w, "INVALID_BODY", http.StatusBadRequest)
		return
	}

	drives, cError := service.GetDriveAccess(groups)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(drives)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
