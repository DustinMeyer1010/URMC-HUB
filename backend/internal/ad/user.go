package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func SearchAllUsers(searchValue string) (matches []models.UserSimpleInfo, err error) {

	matches = make([]models.UserSimpleInfo, 0)

	conn, err := connectToLDAP()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	defer conn.Unbind()

	filter := fmt.Sprintf("(&(objectCategory=user)(|(anr=%s)(URID=%s)))", searchValue, searchValue)

	ldapConfig := SearchConfig(
		filter,
		"cn",
		"name",
		"sAMAccountName",
		"distinguishedName",
		"uid",
		"mail",
		"urid",
	)

	results, err := ldapConfig.Search(conn)

	if results == nil || err != nil {
		return
	}

	for _, entry := range results.Entries {
		var user models.UserSimpleInfo
		user.FillAttributes(entry)
		matches = append(matches, user)
	}

	return
}

func PullUserInformation(searchValue string) (models.UserFullInfo, error) {

	var user models.UserFullInfo

	catagory := "user"
	attribute := "SAMAccountName"

	results, err := SearchAllByCategory(
		catagory,
		attribute,
		searchValue,
		"cn",
		"name",
		"sAMAccountName",
		"distinguishedName",
		"uid",
		"mail",
		"urid",
		"telephoneNumber",
		"department",
		"title",
		"pwdlastset",
		"description",
		"physicalDeliveryOfficeName",
		"givenName",
		"memberOf",
		"URRoleStatus",
		"sn",
	)

	if results == nil {
		return user, fmt.Errorf("username not found")
	}

	if err != nil {
		return user, fmt.Errorf("ldap threw an error: %s", err)
	}

	foundUser := results.Entries[0]

	user.FillAttributes(foundUser)

	return user, err
}
