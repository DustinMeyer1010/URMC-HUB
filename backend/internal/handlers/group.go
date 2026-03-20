package handlers

// NOTE: This file will replace group.go just contains all the new functions

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/parser"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/LostProgrammer1010/URMC-HUB/internal/utils"
)

// HTTP GET requests to retrieve specific LDAP group attributes.
// It expects a 'dn' query parameter for the target object and an optional 'attributes'
// comma-separated list to filter the returned fields.
func GetGroup(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")
	attributes := parser.QueryArray(query.Get("attributes"))

	data, _ := service.GetGroup(dn, attributes...)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// HTTP GET requests to retrieve specific LDAP group all attrubutes
// It expects a 'dn' query parameter for the target object
func GetGroupAvaiableAttributes(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")

	data, _ := service.GetGroupAvaiableAttributes(dn)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// HTTP GET request to retrieve specific members of ldap group.
// It will expect a dn, start, end query. This will return the members
// within the range of start and end up to 1500.
func GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")
	startParm := query.Get("start")
	endParm := query.Get("end")

	start, end := utils.ExtractStartEndRange(startParm, endParm)

	data, cError := service.GetGroupMembers(dn, start, end)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// TODO: Add a user to group
func UsersAddTOGroup(w http.ResponseWriter, r *http.Request) {

}

// TODO: Remove user from group
func UsersRemoveFromGroup(w http.ResponseWriter, r *http.Request) {

}

/*
Refactor: This will have to have a new function create to get all the members of a group.
func GetAllMembersExcel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	group := vars["group"]

	f := service.GetExcelOfAllMembers(group)

	fmt.Println(group)
	if f == nil {
		http.Error(w, "EXCEL_GENERATION_FAILED", http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)

	if err := f.Write(buf); err != nil {
		http.Error(w, "EXCEL_GENERATION_FAILED", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s_Members.xlsx"`, group))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(buf.Bytes())))
	w.Write(buf.Bytes())
}
*/
