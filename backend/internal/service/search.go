package service

import (
	"net/http"
	"net/url"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/gorilla/mux"
)

// Deprecated: No longer needed as search are query based
func getSearchValue(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	searchValue, err := url.QueryUnescape(searchValue)

	if err != nil {
		cError := errs.INVALID_BODY.NewMessage("INVALID SEARCH PARAMETER")
		return "", &cError
	}

	return searchValue, nil

}

// Deprecated: Replace with SearchAllGroupsNew
func SearchAllGroups(r *http.Request) ([]models.GroupSimpleInfo, error) {

	searchValue, cError := getSearchValue(r)

	if cError != nil {
		return []models.GroupSimpleInfo{}, cError
	}

	return ad.SearchAllGroups(searchValue)
}

// Deprecated: Replaces with SearchAll
func AllSearch(r *http.Request) (models.AllResults, error) {
	searchValue, err := getSearchValue(r)

	if err != nil {
		return models.AllResults{}, nil
	}

	return ad.AllSearch(searchValue)
}
