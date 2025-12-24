package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

// Request handler for getting the information about a user
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

// Request handler for getting the lockout information for a user
func LockOutStatus(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)

	vars := mux.Vars(r)
	username := vars["username"]

	matches := ad.LockoutInfoData(username)

	jsonData, _ := json.Marshal(matches)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Request handler for removing a group from a user
func RemoveGroup(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)

	vars := mux.Vars(r)
	username := vars["username"]

	var userModify models.UserModifyMemberOf

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if jsonError := decoder.Decode(&userModify); jsonError != nil {
		cError := customError.INVALID_BODY.NewMessage("INVALID GROUPS IN BODY")
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	modifyResults, cError := ad.RemoveGroup(username, userModify.Groups)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	res, _ := json.Marshal(modifyResults)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// Request handler for adding a group to a user
func AddGroup(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)

	vars := mux.Vars(r)
	username := vars["username"]

	var userModify models.UserModifyMemberOf

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if jsonError := decoder.Decode(&userModify); jsonError != nil {
		cError := customError.INVALID_BODY.NewMessage("INVALID GROUPS IN BODY")
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	modifyResults, cError := ad.AddGroup(username, userModify.Groups)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	res, _ := json.Marshal(modifyResults)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

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

func GetMemberOf(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Memberof")
	logger.LogRequestInfo(r.Method, r.URL.Path)
	vars := mux.Vars(r)
	username := vars["username"]

	groups, cError := service.GetMemberOf(username)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(groups)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
