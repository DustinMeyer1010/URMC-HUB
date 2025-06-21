package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

func SearchAllComputers(searchValue string) (matches []models.ComputerSimpleInfo, err error) {
	matches = make([]models.ComputerSimpleInfo, 0)

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
		fmt.Sprintf("(&(objectCategory=computer)(name=%s*))", searchValue),
		[]string{"name", "operatingSystem", "distinguishedName"},
		nil,
	)

	results, err := l.Search(searchRequest)

	if results == nil || err != nil {
		return
	}
	fmt.Println(len(results.Entries))
	for _, entry := range results.Entries {
		matches = append(matches, models.ComputerSimpleInfo{
			Type:            "Computer",
			Name:            entry.GetAttributeValue("name"),
			OU:              entry.GetAttributeValue("distinguishedName"),
			OperatingSystem: entry.GetAttributeValue("operatingSystem"),
		})
	}

	return

}
