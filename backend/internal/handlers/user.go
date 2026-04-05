package handlers

// NOTE: This file will replace user.go just contains all the new functions

import (
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/LostProgrammer1010/URMC-HUB/internal/parser"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/LostProgrammer1010/URMC-HUB/internal/utils"
)

// HTTP GET requests to retrieve specific LDAP user attributes.
// It expects a 'dn' query parameter for the target object and an optional 'attributes'
// comma-separated list to filter the returned fields.
func GetUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dn := query.Get("dn")
	attributes := parser.QueryArray(query.Get("attributes"))

	data, cError := service.GetUser(dn, attributes...)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
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

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
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

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
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
	var cError error = nil

	fmt.Println(len(attributes))

	if len(attributes) == 0 {
		data, cError = service.GetUserGroups(dn, "samaccountname", "information", "dn", "description", "cn")
	} else {
		data, cError = service.GetUserGroups(dn, attributes...)
	}

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
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

	var userModify models.UserModifyRequest

	// TODO: Handle this error
	err := utils.GatherBodyDetails(r.Body, &userModify)

	fmt.Println(err)

	// NOTE: This should not return and error because all modify request will
	// have the reason on why or why not it was added
	data, _ := service.GroupAddToUser(dn, userModify.GroupDN)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
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

/*
func BulkUserSearch(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)

	var values []string = []string{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if jsonError := decoder.Decode(&values); jsonError != nil {
		http.Error(w, "INVALID_BODY", http.StatusBadRequest)
		return
	}

	f := service.BulkUserSearchValues(values)

	buf := new(bytes.Buffer)
	if f == nil {
		http.Error(w, "EXCEL_GENERATION_FAILED", http.StatusInternalServerError)
		return
	}

	if err := f.Write(buf); err != nil {
		http.Error(w, "EXCEL_GENERATION_FAILED", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", `attachment; filename="example.xlsx"`)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(buf.Bytes())))
	w.Write(buf.Bytes())
}



// Search for all the user in the file and pull simple information
func BulkUserSearchFile(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)
	r.ParseMultipartForm(1 << 20)

	uploadedfiles, ok := r.MultipartForm.File["file"]

	if !ok {
		http.Error(w, "NO_FILE_PROVIDED", http.StatusBadRequest)
		return
	}

	f := service.BulkUserSearch(uploadedfiles)

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
	w.Header().Set("Content-Disposition", `attachment; filename="example.xlsx"`)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(buf.Bytes())))
	w.Write(buf.Bytes())
}
*/
