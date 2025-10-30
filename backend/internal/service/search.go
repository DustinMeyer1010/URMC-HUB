package service

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/gorilla/mux"
)

func getSearchValue(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	searchValue, ok := vars["searchValue"]

	if !ok {
		return "", fmt.Errorf("no search value provided")
	}

	searchValue, err := url.QueryUnescape(searchValue)

	if err != nil {
		return "", fmt.Errorf("failed to escape query")
	}

	return searchValue, nil

}

func SearchAllGroups(r *http.Request) ([]models.GroupSimpleInfo, error) {

	searchValue, err := getSearchValue(r)

	if err != nil {
		return []models.GroupSimpleInfo{}, err
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
