package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

func AddUsersToGroup(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)
	vars := mux.Vars(r)
	group := vars["group"]

	var modify models.ModifyMembers

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if jsonError := decoder.Decode(&modify); jsonError != nil {
		cError := customError.INVALID_BODY.NewMessage("INVALID GROUPS ARRAY")
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	results, cError := ad.AddUsersToGroup(group, modify.Members)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(results)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func RemoveUsersFromGroup(w http.ResponseWriter, r *http.Request) {
	logger.Infof("%s %s", r.Method, r.URL)
	vars := mux.Vars(r)
	group := vars["group"]

	var modify models.ModifyMembers

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if jsonError := decoder.Decode(&modify); jsonError != nil {
		cError := customError.INVALID_BODY.NewMessage("INVALID GROUPS ARRAY")
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	results, cError := ad.RemoveUsersFromGroup(group, modify.Members)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(results)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func GetAllMembers(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)
	vars := mux.Vars(r)
	group := vars["group"]

	membersDNs, cError := service.GetAllMembers(group)
	members := ad.PullMembersInformation(membersDNs)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(members)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func PullGroupInfo(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)
	vars := mux.Vars(r)
	group := vars["group"]

	result, cError := ad.PullGroupInfo(group)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
