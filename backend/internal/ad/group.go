package ad

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

// Pull all groups that match the search value returning CN, distinguishedName, description, info for each group
func SearchAllGroups(searchValue string) ([]models.GroupSimpleInfo, error) {
	matches := make([]models.GroupSimpleInfo, 0)

	results, err := SearchAllByCategory(
		"group",
		"cn",
		searchValue,
		"cn",
		"distinguishedName",
		"sAMAccountName",
		"description",
		"info",
	)

	if results == nil || err != nil {
		return matches, err
	}

	for _, entry := range results.Entries {
		matches = append(matches, models.ToGroupSimpleInfo(entry))
	}

	return matches, err
}

func PullGroupInfo(group string) (models.GroupSimpleInfo, error) {
	groupInfo := models.GroupSimpleInfo{}

	results, err := SearchByCategory(
		"group",
		"cn",
		group,
		"cn",
		"distinguishedName",
		"sAMAccountName",
		"description",
		"info")

	if results == nil || err != nil {
		return groupInfo, err
	}

	entry := results.Entries[0]

	return models.ToGroupSimpleInfo(entry), err
}
