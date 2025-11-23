package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

func AllSearch(w http.ResponseWriter, r *http.Request) {

	matches, statusError := service.AllSearch(r)

	if statusError != nil {
		http.Error(w, statusError.ErrorType, http.StatusUnauthorized)
		w.Write([]byte(statusError.Error))
		return
	}

	jsonData, err := json.Marshal(matches)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func UserSearch(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	searchValue, _ = url.QueryUnescape(searchValue)

	userMatches, err := ad.SearchAllUsers(searchValue)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(userMatches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GroupSearch(w http.ResponseWriter, r *http.Request) {

	groupMatches, statusError := service.SearchAllGroups(r)

	if statusError != nil {
		http.Error(w, statusError.ErrorType, http.StatusUnauthorized)
		w.Write([]byte(statusError.Error))
		return
	}

	jsonData, err := json.Marshal(groupMatches)

	if err != nil {
		http.Error(w, "failed to parse response to json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func PrinterSearch(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	printerMatches, err := ad.SearchAllPrinters(searchValue)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(printerMatches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func ComputerSearch(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	searchValue, _ = url.QueryUnescape(searchValue)

	computerMatches, err := ad.SearchAllComputers(searchValue)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(computerMatches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func ShareDriveSearch(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	searchValue := vars["searchValue"]
	searchValue, _ = url.QueryUnescape(searchValue)
	driveMatches, _ := ad.SearchAllDrives(searchValue)

	jsonData, _ := json.Marshal(driveMatches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
