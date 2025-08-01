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

	ldapConfig := SearchConfig(
		[]string{"name", "operatingSystem", "distinguishedName"},
		fmt.Sprintf("(&(objectCategory=computer)(name=%s*))", searchValue),
	)

	searchRequest := ldap.NewSearchRequest(
		ldapConfig.BaseDN,
		ldapConfig.Scope,
		ldapConfig.Deref,
		ldapConfig.SizeLimit,
		ldapConfig.TimeLimit,
		ldapConfig.TypesOnly,
		ldapConfig.Filter,
		ldapConfig.Attribute,
		ldapConfig.Control,
	)

	results, err := l.Search(searchRequest)

	if results == nil || err != nil {
		return
	}
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
