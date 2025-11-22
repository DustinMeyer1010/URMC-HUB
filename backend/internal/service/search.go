package service

import (
	"net/http"
	"net/url"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/gorilla/mux"
)

func getSearchValue(r *http.Request) (string, *models.Error) {
	vars := mux.Vars(r)
	searchValue := vars["searchValue"]

	searchValue, err := url.QueryUnescape(searchValue)

	if err != nil {
		return "", models.NewError(http.StatusBadRequest, "INVALID_SEARCH", err.Error())
	}

	return searchValue, nil

}

func SearchAllGroups(r *http.Request) ([]models.GroupSimpleInfo, *models.Error) {

	searchValue, statusError := getSearchValue(r)

	if statusError != nil {
		return []models.GroupSimpleInfo{}, statusError
	}

	return ad.SearchAllGroups(searchValue)
}

func AllSearch(r *http.Request) (models.AllResults, error) {
	searchValue, err := getSearchValue(r)

	if err != nil {
		return models.AllResults{}, err
	}

	return ad.AllSearch(searchValue)
}
