package ad

import (
	"fmt"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

var LDAP_STRING_REPLACE = strings.NewReplacer(
	`\`, `\5c`,
	`*`, `\2a`,
	`(`, `\28`,
	`)`, `\29`,
	"\x00", `\00`,
)

// Pull all groups that match the search value returning CN, distinguishedName, description, info for each group
func SearchAllGroups(searchValue string) ([]models.GroupSimpleInfo, *customError.Error) {
	matches := make([]models.GroupSimpleInfo, 0)
	searchValue = LDAP_STRING_REPLACE.Replace(searchValue)

	results, ldapError := SearchAllByCategory(
		"group",
		"cn",
		searchValue,
		"cn",
		"distinguishedName",
		"sAMAccountName",
		"description",
		"info",
	)

	if ldapError != nil {
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return matches, &cError
	}

	if results == nil {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO RESULTS FOUND FOR: %s", searchValue))
		return matches, &cError
	}

	for _, entry := range results.Entries {
		matches = append(matches, models.ToGroupSimpleInfo(entry))
	}

	return matches, nil
}

func PullGroupInfo(group string) (models.GroupSimpleInfo, *customError.Error) {
	groupInfo := models.GroupSimpleInfo{}
	group = LDAP_STRING_REPLACE.Replace(group)

	results, ldapError := SearchByCategory(
		"group",
		"cn",
		group,
		"cn",
		"distinguishedName",
		"sAMAccountName",
		"description",
		"info")

	if ldapError != nil {
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return groupInfo, &cError
	}

	if results == nil {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO GROUP FOUND FOR: %s", group))
		return groupInfo, &cError
	}

	if len(results.Entries) > 1 {
		fmt.Println(results.Entries)
	}

	entry := results.Entries[0]

	return models.ToGroupSimpleInfo(entry), nil
}

// TODO: Create a message per users letting the frontend know if that user was or was not added to the group
func AddUsersToGroup(group string, newMembers []string) *customError.Error {

	usersDN, cError := GetUserDN(newMembers)
	if cError != nil {
		return cError
	}
	group = LDAP_STRING_REPLACE.Replace(group)
	groupDN, cError := GetGroupDN(group)

	for _, user := range usersDN {
		cError = ModifyGroupNewMember(groupDN, user)
		if cError != nil {
			// User not added but frontend does not know
			fmt.Println(cError)
			continue
		}
		fmt.Printf("User: %s\n Added to Group\n", user)
	}

	return nil
}

// TODO: Create a message per users letting the frontend know if that user was or was not removed to the group
func RemoveUsersFromGroup(group string, members []string) *customError.Error {
	l, cError := connectToLDAP()
	if cError != nil {
		return cError
	}
	defer l.Close()
	defer l.Unbind()

	var usersDN []string
	if usersDN, cError = GetUserDN(members); cError != nil {
		return cError
	}

	var groupDN string
	if groupDN, cError = GetGroupDN(group); cError != nil {
		return cError
	}

	for _, user := range usersDN {
		cError = ModifyGroupRemoveMember(groupDN, user)
		if cError != nil {
			fmt.Println(cError)
			continue
		}
		fmt.Printf("User: %s\n Removed From Group\n", user)
	}

	return nil
}

func GetGroupsDN(groups []string) ([]string, *customError.Error) {
	groupsDN := make([]string, 0)

	for _, group := range groups {
		DN, cError := GetGroupDN(group)

		// No group found with the given name
		if cError != nil {
			if cError.Type == "LDAP_ERROR" {
				return groupsDN, cError
			}
			continue
		}

		groupsDN = append(groupsDN, DN)
	}

	return groupsDN, nil
}

func GetGroupDN(group string) (string, *customError.Error) {

	l, cError := connectToLDAP()
	if cError != nil {
		return "", cError
	}
	defer l.Close()
	defer l.Unbind()

	group = LDAP_STRING_REPLACE.Replace(group)

	results, ldapError := SearchByCategory("group", "cn", group, "dn")

	if ldapError != nil {
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return "", &cError
	}

	if len(results.Entries) == 0 {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO GROUP FOUND FOR: %s", group))
		return "", &cError
	}

	return results.Entries[0].DN, nil
}

func ModifyGroupNewMember(groupDN, user string) *customError.Error {
	l, cError := connectToLDAP()

	if cError != nil {
		return cError
	}
	defer l.Close()
	defer l.Unbind()

	addRequest := ldap.NewModifyRequest(groupDN, nil)
	addRequest.Add("member", []string{user})
	ldapError := l.Modify(addRequest)

	if ldapError != nil {
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return &cError
	}

	return nil
}

func ModifyGroupRemoveMember(groupDN, user string) *customError.Error {
	l, cError := connectToLDAP()

	if cError != nil {
		return cError
	}
	defer l.Close()
	defer l.Unbind()

	addRequest := ldap.NewModifyRequest(groupDN, nil)
	addRequest.Delete("member", []string{user})
	ldapError := l.Modify(addRequest)

	if ldapError != nil {
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return &cError
	}

	return nil
}
