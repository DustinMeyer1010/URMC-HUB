package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

// Deprecated: Replaced with SearchAll
func AllSearch(w http.ResponseWriter, r *http.Request) {

	logger.LogRequestInfo(r.Method, r.URL.Path)

	matches, cError := service.AllSearch(r)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	jsonData, _ := json.Marshal(matches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

// Deprecated: Replaced with SearchUsers
func UserSearch(w http.ResponseWriter, r *http.Request) {

	logger.LogRequestInfo(r.Method, r.URL.Path)

	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	searchValue, _ = url.QueryUnescape(searchValue)

	userMatches, cError := ad.SearchAllUsers(searchValue)

	if e := errs.IsApiError(cError); e != nil {
		w.WriteHeader(e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	jsonData, _ := json.Marshal(userMatches)

	logger.Info(jsonData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Deprecated: Replaced with SearchGroups
func GroupSearch(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)

	groupMatches, cError := service.SearchAllGroups(r)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	jsonData, _ := json.Marshal(groupMatches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

// Deprecated: Replaced with SearchPrinters
func PrinterSearch(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)

	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	printerMatches, cError := ad.SearchAllPrinters(searchValue)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	jsonData, _ := json.Marshal(printerMatches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

// Deprecated: Replaced with SearchComputer
func ComputerSearch(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)

	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	searchValue, _ = url.QueryUnescape(searchValue)

	computerMatches, cError := ad.SearchAllComputers(searchValue)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	jsonData, _ := json.Marshal(computerMatches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

// Deprecated: Replaced with SearchDrives
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
