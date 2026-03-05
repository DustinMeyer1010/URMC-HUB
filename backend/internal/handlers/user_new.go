package handlers

// NOTE: This file will replace user.go just contains all the new functions

import (
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/parser"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
)

// HTTP GET requests to retrieve specific LDAP user attributes.
// It expects a 'dn' query parameter for the target object and an optional 'attributes'
// comma-separated list to filter the returned fields.
func GetUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")
	attributes := parser.QueryArray(query.Get("attributes"))

	data, cError := service.GetUser(dn, attributes...)

	if cError != nil {
		w.WriteHeader(cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// HTTP GET requests to retrieve specific LDAP user all attrubutes
// It expects a 'dn' query parameter for the target object
func GetUserAvaiableAttributes(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")

	data, cError := service.GetUserAvaiableAttributes(dn)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// HTTP GET request to retrive specific ldap user and all the drives
// that that user has access to. It expects a 'dn' query parameter for the target object
func GetUserDrives(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")
	data, cError := service.GetUserDrives(dn)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// HTTP GET request to retrive specifc ldap user and all the groups
// they have on their account. It expects a 'dn' query parameter for the target object.
// An optional attributes query parameter for the attributes of the groups wanted on
// response
func GetUserGroups(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")
	attributes := parser.QueryArray(query.Get("attributes"))

	data := make([]byte, 0)
	var cError *customError.Error = nil

	if len(attributes) == 0 {
		data, cError = service.GetUserGroups(dn, "samaccountname", "information", "dn", "description", "cn")
	} else {
		data, cError = service.GetUserGroups(dn, attributes...)
	}

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

// HTTP GET request to retrive specifc ldap user password attempt history.
// expects a 'dn'
func GetUserLockoutStatus(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	dn := query.Get("dn")

	data := service.GetUserLockoutStatus(dn)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// HTTP POST request to add specific ldap user from a group
func GroupAddToUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")

	fmt.Println(dn)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Still Being Implemented"))
}

// HTTP DELETE request to remove specific ldap user from a group
func GroupRemoveFromUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")

	fmt.Println(dn)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Still Being Implemented"))
}
