package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

func AllSearch(w http.ResponseWriter, r *http.Request) {

	logger.LogRequestInfo(r.Method, r.URL.Path)

	matches, cError := service.AllSearch(r)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(matches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func UserSearch(w http.ResponseWriter, r *http.Request) {

	logger.LogRequestInfo(r.Method, r.URL.Path)

	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	searchValue, _ = url.QueryUnescape(searchValue)

	userMatches, cError := ad.SearchAllUsers(searchValue)

	if cError != nil {
		w.WriteHeader(cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(userMatches)

	logger.Info(jsonData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GroupSearch(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)

	groupMatches, cError := service.SearchAllGroups(r)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(groupMatches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func PrinterSearch(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)

	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	printerMatches, cError := ad.SearchAllPrinters(searchValue)

	if cError != nil {
		w.WriteHeader(cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(printerMatches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func ComputerSearch(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)

	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	searchValue, _ = url.QueryUnescape(searchValue)

	computerMatches, cError := ad.SearchAllComputers(searchValue)

	if cError != nil {
		w.WriteHeader(cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(computerMatches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func ShareDriveSearch(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)

	vars := mux.Vars(r)
	searchValue := vars["searchValue"]
	searchValue, _ = url.QueryUnescape(searchValue)
	driveMatches, _ := ad.SearchAllDrives(searchValue)

	jsonData, _ := json.Marshal(driveMatches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
