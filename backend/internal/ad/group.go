package ad

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

// Pull all groups that match the search value returning CN, distinguishedName, description, info for each group
func SearchAllGroups(searchValue string) ([]models.GroupSimpleInfo, *models.Error) {
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
		return matches, models.NewError(http.StatusNotFound, "NOT_FOUND", "No results found")
	}

	for _, entry := range results.Entries {
		matches = append(matches, models.ToGroupSimpleInfo(entry))
	}

	return matches, nil
}

func PullGroupInfo(group string) (models.GroupSimpleInfo, error) {
	groupInfo := models.GroupSimpleInfo{}

	r := strings.NewReplacer(
		`\`, `\5c`,
		`*`, `\2a`,
		`(`, `\28`,
		`)`, `\29`,
		"\x00", `\00`,
	)

	group = r.Replace(group)

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
		fmt.Println()
		fmt.Println(group)
		fmt.Println(err)
		fmt.Println()
		return groupInfo, err
	}

	if len(results.Entries) > 1 {
		fmt.Println(results.Entries)
	}

	entry := results.Entries[0]

	return models.ToGroupSimpleInfo(entry), err
}

func GetGroupsDN(groups []string) ([]string, error) {

	var groupsDN []string

	for _, group := range groups {
		DN, err := GetGroupDN(group)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		groupsDN = append(groupsDN, DN)
	}

	return groupsDN, nil
}

func AddUsersToGroup(group string, newMembers []string) error {

	l, err := connectToLDAP()
	if err != nil {
		return err
	}
	defer l.Close()
	defer l.Unbind()

	var usersDN []string
	if usersDN, err = GetUserDN(newMembers, l); err != nil {
		return err
	}

	var groupDN string
	if groupDN, err = GetGroupDN(group); err != nil {
		return err
	}

	for _, user := range usersDN {
		addRequest := ldap.NewModifyRequest(groupDN, nil)
		addRequest.Add("member", []string{user})
		groupAddError := l.Modify(addRequest)
		if groupAddError != nil {
			fmt.Printf("\nError: %sUser: %s\n", groupAddError.Error(), user)
			continue
		}
		fmt.Printf("User: %s\n Added to Group\n", user)
	}

	return nil
}

func RemoveUsersFromGroup(group string, members []string) error {
	l, err := connectToLDAP()
	if err != nil {
		return err
	}
	defer l.Close()
	defer l.Unbind()

	var usersDN []string
	if usersDN, err = GetUserDN(members, l); err != nil {
		return err
	}

	var groupDN string
	if groupDN, err = GetGroupDN(group); err != nil {
		return err
	}

	for _, user := range usersDN {
		addRequest := ldap.NewModifyRequest(groupDN, nil)
		addRequest.Delete("member", []string{user})
		groupAddError := l.Modify(addRequest)
		if groupAddError != nil {
			fmt.Printf("\nError: %sUser: %s\n", groupAddError.Error(), user)
			continue
		}
		fmt.Printf("User: %s\n Removed From Group\n", user)
	}

	return nil
}

func GetGroupDN(group string) (string, error) {

	l, err := connectToLDAP()
	if err != nil {
		return "", err
	}
	defer l.Close()
	defer l.Unbind()

	config := SearchConfig(fmt.Sprintf("(&(objectClass=group)(cn=%s))", group), "dn")

	results, err := config.Search(l)

	if err != nil {
		return "", err
	}
	if len(results.Entries) == 0 {
		return "", fmt.Errorf("No Entries Found for Group: %s", group)
	}

	return results.Entries[0].DN, nil
}
