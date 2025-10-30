package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/gorilla/mux"
)

func PullUserInformation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := ad.PullUserInformation(username)

	if err != nil {
		fmt.Println(err)
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

}

func AddGroup(w http.ResponseWriter, r *http.Request) {

}
