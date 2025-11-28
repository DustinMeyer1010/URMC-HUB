package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/gorilla/mux"
)

func AddUsersToGroup(w http.ResponseWriter, r *http.Request) {
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

	cError := ad.AddUsersToGroup(group, modify.Members)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}
}

func RemoveUsersFromGroup(w http.ResponseWriter, r *http.Request) {
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

	cError := ad.RemoveUsersFromGroup(group, modify.Members)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

}
