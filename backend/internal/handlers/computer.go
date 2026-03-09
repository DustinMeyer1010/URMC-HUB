package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

// Deprecated: To be replaced by GetComputer
func ComputerInfo(w http.ResponseWriter, r *http.Request) {

	logger.LogRequestInfo(r.Method, r.URL.Path)

	vars := mux.Vars(r)
	computer := vars["computer"]

	computerResponse, cError := service.ComputerInfo(computer)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	jsonData, _ := json.Marshal(computerResponse)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
