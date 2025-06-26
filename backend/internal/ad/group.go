package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

func SearchAllGroups(searchValue string) (matches []models.GroupSimpleInfo, err error) {
	matches = make([]models.GroupSimpleInfo, 0)

	l, err := connectToLDAP()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()
	defer l.Unbind()

	searchRequest := ldap.NewSearchRequest(
		"DC=URMC-sh,DC=rochester,DC=edu",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectCategory=group)((anr=%s)))", searchValue),
		[]string{"cn", "distinguishedName", "sAMAccountName", "description", "info"},
		nil,
	)

	results, err := l.Search(searchRequest)
	fmt.Println(err)

	if results == nil || err != nil {
		return
	}

	for _, entry := range results.Entries {
		matches = append(matches,
			models.GroupSimpleInfo{
				Type:        "Group",
				Name:        entry.GetAttributeValue("sAMAccountName"),
				OU:          entry.GetAttributeValue("distinguishedName"),
				Description: entry.GetAttributeValue("description"),
				Information: entry.GetAttributeValue("info"),
			})

	}

	return
}
