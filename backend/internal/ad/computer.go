package ad

import (
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
