package handlers

// NOTE: This file will replace user.go just contains all the new functions

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/parser"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

// HTTP GET requests to retrieve specific LDAP user attributes.
// It expects a 'dn' query parameter for the target object and an optional 'attributes'
// comma-separated list to filter the returned fields.
func GetUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")
	attributes := parser.QueryArray(query.Get("attributes"))

	jsonData, cError := service.GetUser(dn, attributes...)

	if cError != nil {
		w.WriteHeader(cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// HTTP GET requests to retrieve specific LDAP user all attrubutes
// It expects a 'dn' query parameter for the target object
func GetUserAvaiableAttributes(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")

	jsonData, _ := service.GetUserAvaiableAttributes(dn)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// Deprecated: Going to be replaced by GetUser
func PullUserInformation(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)
	vars := mux.Vars(r)
	username := vars["username"]

	user, cError := ad.PullUserInformation(username)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(user)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// TODO: Create desctiption for function
func GetUserDrives(w http.ResponseWriter, r *http.Request) {
	// TODO: Get all the drives the user will have access to
}

// TODO: Create desctiption for function
func GetUserGroups(w http.ResponseWriter, r *http.Request) {
	// TODO: Get all groups on user account
}

// TODO: Create desctiption for function
func GetUserLockoutStatus(w http.ResponseWriter, r *http.Request) {
	// TODO: Get Lockout Info for user account
}
