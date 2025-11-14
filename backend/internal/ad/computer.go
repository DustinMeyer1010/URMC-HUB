package ad

import (
	"fmt"

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

func PullComputerInformation(comptuer string) (models.ComputerSimpleInfo, error) {

	results, err := SearchByCategory(
		"computer",
		"name",
		comptuer,
		"name",
		"operatingSystem",
		"distinguishedName",
	)

	if err != nil {
		return models.ComputerSimpleInfo{}, err
	}

	if len(results.Entries) == 0 {
		return models.ComputerSimpleInfo{}, fmt.Errorf("%s", "No Results Found")
	}

	return models.ToComputerSimpleInfo(results.Entries[0]), nil

}
