package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

// Pull all groups that match the search value returning CN, distinguishedName, description, info for each group
func SearchAllGroups(searchValue string) (matches []models.GroupSimpleInfo, err error) {
	matches = make([]models.GroupSimpleInfo, 0)

	conn, err := connectToLDAP()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	defer conn.Unbind()

	ldapConfig := SearchConfig(
		[]string{"cn", "distinguishedName", "sAMAccountName", "description", "info"},
		fmt.Sprintf("(&(objectCategory=group)(cn=%s*))", searchValue),
	)

	results, err := ldapConfig.Search(conn)

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

func PullGroupInfo(group string) (*models.GroupSimpleInfo, error) {

	conn, err := connectToLDAP()

	if err != nil {
		return nil, err
	}

	defer conn.Close()
	defer conn.Unbind()

	ldapConfig := SearchConfig(
		[]string{"cn", "distinguishedName", "sAMAccountName", "description", "info"},
		fmt.Sprintf("(&(objectCategory=group)(|(cn=%s))", group),
	)

	results, err := ldapConfig.Search(conn)

	if results == nil || err != nil {
		return nil, err
	}

	entry := results.Entries[0]

	return &models.GroupSimpleInfo{
		Name:        entry.GetAttributeValue("cn"),
		Description: entry.GetAttributeValue("description"),
		Information: entry.GetAttributeValue("info"),
		OU:          entry.GetAttributeValue("distinguishedName"),
	}, nil
}
