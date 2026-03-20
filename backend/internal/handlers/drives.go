package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
)

// Deprecated: Replace with GetUserDrives
func DriveAccess(w http.ResponseWriter, r *http.Request) {

	logger.LogRequestInfo(r.Method, r.URL.Path)

	var groups []string

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&groups); err != nil {
		http.Error(w, "INVALID_BODY", http.StatusBadRequest)
		return
	}

	drives, cError := service.GetDriveAccess(groups)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	jsonData, _ := json.Marshal(drives)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func DriveInfo(w http.ResponseWriter, r *http.Request) {

	logger.LogRequestInfo(r.Method, r.URL.Path)
	query := r.URL.Query()

	drive := query.Get("drive")

	fmt.Println(drive)

	driveInfo := ad.GetShareDriveGroups(drive)

	jsonData, _ := json.Marshal(driveInfo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
