package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

// Pull all groups that match the search value returning CN, distinguishedName, description, info for each group
func SearchAllGroups(searchValue string) (matches []models.GroupSimpleInfo, err error) {
	matches = make([]models.GroupSimpleInfo, 0)

	l, err := connectToLDAP()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()
	defer l.Unbind()

	ldapConfig := SearchConfig(
		[]string{"cn", "distinguishedName", "sAMAccountName", "description", "info"},
		fmt.Sprintf("(&(objectCategory=user)(|(anr=%s)(URID=%s)))", searchValue, searchValue),
	)

	results, err := ldapConfig.Search(l)

	if results == nil || err != nil {
		return
	}

	for _, entry := range results.Entries {
		matches = append(matches,
			models.GroupSimpleInfo{
				Name:        entry.GetAttributeValue("sAMAccountName"),
				OU:          entry.GetAttributeValue("distinguishedName"),
				Description: entry.GetAttributeValue("description"),
				Information: entry.GetAttributeValue("info"),
			})

	}

	return
}
