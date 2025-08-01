package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
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

	ldapConfig := SearchConfig(
		[]string{"cn", "name", "sAMAccountName", "distinguishedName", "uid", "mail"},
		fmt.Sprintf("(&(objectCategory=user)(|(anr=%s)(URID=%s)))", searchValue, searchValue),
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

	results, err := conn.Search(searchRequest)

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

func PullUserInformation(username string) (models.UserFullInfo, error) {

	var user models.UserFullInfo

	conn, err := connectToLDAP()

	if err != nil {
		return user, err
	}

	defer conn.Close()
	defer conn.Unbind()

	ldapConfig := SearchConfig(
		[]string{
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
			"sn",
		},
		fmt.Sprintf("(&(objectCategory=user)(|(SAMAccountName=%s)))", username),
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

	results, err := conn.Search(searchRequest)

	if results == nil {
		return user, fmt.Errorf("username not found")
	}

	if err != nil {
		return user, fmt.Errorf("ldap threw an error: %s", err)
	}

	foundUser := results.Entries[0]

	user.FillAttributes(foundUser)

	return user, nil
}
