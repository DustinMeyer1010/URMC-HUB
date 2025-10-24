package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/gorilla/mux"
)

func AllSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	decodedSearch, err := url.QueryUnescape(searchValue)

	if err != nil {
		http.Error(w, "invalid search value", http.StatusBadRequest)
		return
	}

	matches, err := ad.AllSearch(decodedSearch)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	searchValue, _ = url.QueryUnescape(searchValue)

	groupMatches, err := ad.SearchAllGroups(searchValue)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(groupMatches)

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
