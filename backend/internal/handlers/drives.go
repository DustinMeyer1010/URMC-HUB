package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
)

func DriveAccess(w http.ResponseWriter, r *http.Request) {

	var groups []string

	json.NewDecoder(r.Body).Decode(&groups)

	drives := service.GetDriveAccess(groups)

	data, err := json.Marshal(drives)

	if err != nil {
		http.Error(w, "Failed to write output to jsaon", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
