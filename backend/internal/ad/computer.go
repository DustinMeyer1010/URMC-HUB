package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func SearchAllComputers(searchValue string) (matches []models.ComputerSimpleInfo, err error) {
	matches = make([]models.ComputerSimpleInfo, 0)

	conn, err := connectToLDAP()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	defer conn.Unbind()

	ldapConfig := SearchConfig(
		[]string{"name", "operatingSystem", "distinguishedName"},
		fmt.Sprintf("(&(objectCategory=computer)(name=%s*))", searchValue),
	)

	results, err := ldapConfig.Search(conn)

	if results == nil || err != nil {
		return
	}
	for _, entry := range results.Entries {
		matches = append(matches, models.ComputerSimpleInfo{
			Name:            entry.GetAttributeValue("name"),
			OU:              entry.GetAttributeValue("distinguishedName"),
			OperatingSystem: entry.GetAttributeValue("operatingSystem"),
		})
	}

	return

}
