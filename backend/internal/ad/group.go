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
func AddUsersToGroup(group string, newMembers []string) (map[string]models.GroupModifyResults, *customError.Error) {

	var results map[string]models.GroupModifyResults = make(map[string]models.GroupModifyResults)

	usersDN, cError := GetUsersDN(newMembers)
	if cError != nil {
		return results, cError
	}
	group = LDAP_STRING_REPLACE.Replace(group)
	groupDN, cError := GetGroupDN(group)

	if cError != nil {
		return results, cError
	}

	for user, userDN := range usersDN {
		cError = ModifyGroupNewMember(groupDN, userDN)
		if cError != nil {
			results[user] = models.GroupModifyResults{
				Group:      group,
				Message:    cError.Msg,
				Successful: false,
			}
			continue
		}
		results[user] = models.GroupModifyResults{Group: group, Message: "Added to Group", Successful: true}
	}

	return results, nil
}

// TODO: Create a message per users letting the frontend know if that user was or was not removed to the group
func RemoveUsersFromGroup(group string, members []string) (map[string]models.GroupModifyResults, *customError.Error) {

	var cError *customError.Error
	var results map[string]models.GroupModifyResults = make(map[string]models.GroupModifyResults)

	var usersDN map[string]string
	if usersDN, cError = GetUsersDN(members); cError != nil {
		return results, cError
	}

	var groupDN string
	if groupDN, cError = GetGroupDN(group); cError != nil {
		return results, cError
	}

	for user, userDN := range usersDN {
		cError = ModifyGroupRemoveMember(groupDN, userDN)
		if cError != nil {
			results[user] = models.GroupModifyResults{
				Group:      group,
				Message:    cError.Msg,
				Successful: false,
			}
			continue
		}
		results[user] = models.GroupModifyResults{Group: group, Message: "Removed from Group", Successful: true}
	}

	return results, nil
}

func GetGroupsDN(groups []string) (map[string]string, *customError.Error) {
	groupsDN := make(map[string]string, 0)

	for _, group := range groups {
		DN, cError := GetGroupDN(group)

		// No group found with the given name
		if cError != nil {
			if cError.Type == "LDAP_ERROR" {
				return groupsDN, cError
			}
			continue
		}

		groupsDN[group] = DN
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
