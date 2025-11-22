package ad

import (
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func SearchAllComputers(searchValue string) ([]models.ComputerSimpleInfo, error) {
	matches := make([]models.ComputerSimpleInfo, 0)

	results, err := SearchAllByCategory(
		"computer",
		"name",
		searchValue,
		"name",
		"operatingSystem",
		"distinguishedName",
	)

	if results == nil || err != nil {
		return matches, err
	}

	for _, entry := range results.Entries {
		matches = append(matches, models.ToComputerSimpleInfo(entry))
	}

	return matches, err
}

func PullComputerInformation(computer string) (models.ComputerSimpleInfo, *models.Error) {

	results, err := SearchByCategory(
		"computer",
		"name",
		computer,
		"name",
		"operatingSystem",
		"distinguishedName",
	)

	if err != nil {
		return models.ComputerSimpleInfo{},
			models.NewError(http.StatusInternalServerError, "LDAP_ERROR", err.Error())
	}

	if len(results.Entries) == 0 {
		return models.ComputerSimpleInfo{},
			models.NewError(http.StatusNotFound, "NOT_FOUND", fmt.Sprintf("No computer found for %s", computer))
	}

	return models.ToComputerSimpleInfo(results.Entries[0]), nil

}
