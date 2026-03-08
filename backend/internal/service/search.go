package service

import (
	"net/http"
	"net/url"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/gorilla/mux"
)

// Deprecated: No longer needed as search are query based
func getSearchValue(r *http.Request) (string, *customError.Error) {
	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	searchValue, err := url.QueryUnescape(searchValue)

	if err != nil {
		cError := customError.INVALID_BODY.NewMessage("INVALID SEARCH PARAMETER")
		return "", &cError
	}

	return searchValue, nil

}

// Deprecated: Replace with SearchAllGroupsNew
func SearchAllGroups(r *http.Request) ([]models.GroupSimpleInfo, *customError.Error) {

	searchValue, cError := getSearchValue(r)

	if cError != nil {
		return []models.GroupSimpleInfo{}, cError
	}

	return ad.SearchAllGroups(searchValue)
}

// Deprecated: Replaces with SearchAll
func AllSearch(r *http.Request) (models.AllResults, *customError.Error) {
	searchValue, err := getSearchValue(r)

	if err != nil {
		return models.AllResults{}, nil
	}

	return ad.AllSearch(searchValue)
}
