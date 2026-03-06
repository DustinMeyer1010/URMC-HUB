package handlers

// NOTE: This file will replace group.go just contains all the new functions

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/parser"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
)

// HTTP GET requests to retrieve specific LDAP group attributes.
// It expects a 'dn' query parameter for the target object and an optional 'attributes'
// comma-separated list to filter the returned fields.
func GetGroup(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")
	attributes := parser.QueryArray(query.Get("attributes"))

	jsonData, _ := service.GetGroup(dn, attributes...)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// HTTP GET requests to retrieve specific LDAP group all attrubutes
// It expects a 'dn' query parameter for the target object
func GetGroupAvaiableAttributes(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")

	jsonData, _ := service.GetGroupAvaiableAttributes(dn)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// TODO: Create descrtipiion of the function
func GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	// TODO: This endpoint should get all members the group based
	// on the search parameter
}
