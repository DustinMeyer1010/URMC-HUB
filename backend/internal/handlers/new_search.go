package handlers

import (
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/parser"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
)

func SearchAll(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	searchValue := query.Get("value")

	if searchValue == "" {
		http.Error(w, "INVALID_SEARCH_VALUE", http.StatusBadRequest)
		return
	}

	data := service.SearchAll(searchValue)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	searchValue := query.Get("value")
	attributes := parser.QueryArray(query.Get("attributes"))

	if len(attributes) == 0 {
		attributes = []string{"*"}
	}

	data, cError := service.SearchAllUsers(searchValue, attributes...)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func SearchGroups(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	searchValue := query.Get("value")
	attributes := parser.QueryArray(query.Get("attributes"))

	if len(attributes) == 0 {
		attributes = []string{"*"}
	}

	data, cError := service.SearchAllGroupsNew(searchValue, attributes...)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func SearchComputer(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	searchValue := query.Get("value")
	attributes := parser.QueryArray(query.Get("attributes"))

	if len(attributes) == 0 {
		attributes = []string{"*"}
	}

	data, cError := service.SearchAllComputers(searchValue, attributes...)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func SearchDrives(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	searchValue := query.Get("value")

	data, cError := service.SearchAllDrives(searchValue)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func SearchPrinters(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	searchValue := query.Get("value")

	data, cError := service.SearchAllPrinters(searchValue)

	if e := errs.IsApiError(cError); e != nil {
		http.Error(w, e.Type, e.StatusCode)
		w.Write([]byte(e.Msg))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
