package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

func PullUserInformation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := ad.PullUserInformation(username)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed"))
		return
	}

	jsonData, _ := json.Marshal(user)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func LockOutStatus(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]

	matches := ad.LockoutInfoData(username)

	jsonData, _ := json.Marshal(matches)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func RemoveGroup(w http.ResponseWriter, r *http.Request) {

	var groupModify models.GroupModify

	err := json.NewDecoder(r.Body).Decode(&groupModify)

	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	modifyResults, err := ad.RemoveGroup(groupModify.Users, groupModify.Groups)

	if err != nil {
		http.Error(w, "Unable to remove group", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(modifyResults)

	if err != nil {
		http.Error(w, "Failed to parse into JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func AddGroup(w http.ResponseWriter, r *http.Request) {
	var groupModify models.GroupModify

	err := json.NewDecoder(r.Body).Decode(&groupModify)

	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	modifyResults, err := ad.AddGroup(groupModify.Users, groupModify.Groups)

	if err != nil {
		http.Error(w, "Unable to remove group", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(modifyResults)

	if err != nil {
		http.Error(w, "Failed to parse into JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}

func BulkUserSearchFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	r.ParseMultipartForm(1 << 20)

	uploadedfiles, ok := r.MultipartForm.File["file"]

	if !ok {
		http.Error(w, "No Files Provided", http.StatusBadRequest)
		return
	}

	f := service.BulkUserSearch(uploadedfiles)

	buf := new(bytes.Buffer)
	if err := f.Write(buf); err != nil {
		http.Error(w, "Failed to generate Excel file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", `attachment; filename="example.xlsx"`)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(buf.Bytes())))
	w.Write(buf.Bytes())
}
