package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/parser"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

// Deprecated: To be replaced by GetComputer
func ComputerInfo(w http.ResponseWriter, r *http.Request) {

	logger.LogRequestInfo(r.Method, r.URL.Path)

	vars := mux.Vars(r)
	computer := vars["computer"]

	computerResponse, cError := service.ComputerInfo(computer)

	if cError != nil {
		logger.Errorf("%s %s", cError.Type, cError.Msg)
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(computerResponse)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// HTTP GET requests to retrieve specific LDAP computer attributes.
// It expects a 'dn' query parameter for the target object and an optional 'attributes'
// comma-separated list to filter the returned fields.
func GetComputer(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")
	attributes := parser.QueryArray(query.Get("attributes"))

	jsonData, _ := service.GetComputer(dn, attributes...)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// HTTP GET requests to retrieve specific LDAP computer all attrubutes
// It expects a 'dn' query parameter for the target object
func GetComputerAvaiableAttributes(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")

	jsonData, _ := service.GetComputerAvaiableAttributes(dn)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
