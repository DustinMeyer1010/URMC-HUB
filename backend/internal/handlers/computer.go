package handlers

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/parser"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
)

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

func PingComputer(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	computerName := query.Get("name")

	data, err := service.PingComputer(computerName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		w.Write(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}
