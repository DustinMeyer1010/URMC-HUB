package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

// Deprecated: These will fall under edit a person account. Using GroupAddToUser
func AddUsersToGroup(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)
	vars := mux.Vars(r)
	group := vars["group"]

	var modify models.ModifyMembers

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if jsonError := decoder.Decode(&modify); jsonError != nil {
		cError := errs.INVALID_BODY.NewMessage("INVALID GROUPS ARRAY")
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	results, cError := ad.AddUsersToGroup(group, modify.Members)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	jsonData, _ := json.Marshal(results)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Deprecated: These will fall under edit a person account. Using GroupRemoveFromUser
func RemoveUsersFromGroup(w http.ResponseWriter, r *http.Request) {
	logger.Infof("%s %s", r.Method, r.URL)
	vars := mux.Vars(r)
	group := vars["group"]

	var modify models.ModifyMembers

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if jsonError := decoder.Decode(&modify); jsonError != nil {
		cError := errs.INVALID_BODY.NewMessage("INVALID GROUPS ARRAY")
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	results, cError := ad.RemoveUsersFromGroup(group, modify.Members)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	jsonData, _ := json.Marshal(results)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

// Deprecated: Replace with GetGroups
func PullGroupInfo(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)
	vars := mux.Vars(r)
	group := vars["group"]

	result, cError := ad.PullGroupInfo(group)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	jsonData, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Deprecated: Replaced with GetGroupMemebers
func GetAllMembers(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)
	vars := mux.Vars(r)
	group := vars["group"]

	membersDNs, cError := service.GetAllMembers(group)
	members := ad.PullMembersInformation(membersDNs)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	jsonData, _ := json.Marshal(members)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

// Refactor: This will have to have a new function create to get all the members of a group.
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
