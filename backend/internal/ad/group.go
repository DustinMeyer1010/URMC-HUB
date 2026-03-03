package ad

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
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

var RATE_LIMIT = 100

// Pull all groups that match the search value returning CN, distinguishedName, description, info for each group
func SearchAllGroups(searchValue string, attributes ...string) ([]models.GroupSimpleInfo, *customError.Error) {
	groups := []models.GroupSimpleInfo{}

	/*
		"cn",
		"distinguishedName",
		"sAMAccountName",
		"description",
		"info",
	*/

	searchConfig := DefaultSearchConfig().
		SetFilter(fmt.Sprintf("(&(objectCategory=group)(anr=%s*))", searchValue)).
		SetAttributes(attributes)

	searchResults, ldapError := searchConfig.Search()

	if ldapError != nil {
		logger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return groups, &cError
	}

	if searchResults == nil {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO RESULTS FOUND FOR: %s", searchValue))
		return groups, &cError
	}

	for _, entry := range searchResults.Entries {
		groups = append(groups, models.ToGroupSimpleInfo(entry))
	}

	return groups, nil
}

// Performs an LDAP search for a specific group by its Distinguished Name.
// It returns a mapped collection of the requested attributes or a custom error if the
// group is not found or the search fails.
func LookupGroup(groupDN string, attributes ...string) (map[string][]string, *customError.Error) {

	searchConfig := GroupSearchConfig().
		SetBaseDN(groupDN).
		SetAttributes(attributes)

	searchResults, ldapError := searchConfig.Search()

	if ldapError != nil {
		logger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return map[string][]string{}, &cError
	}

	if searchResults == nil || len(searchResults.Entries) == 0 {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO GROUP FOUND FOR: %s", groupDN))
		logger.Errorf("%v", cError)
		logger.Debug(searchResults.Entries)
		return map[string][]string{}, &cError
	}

	entry := searchResults.Entries[0]

	attrs := createAttributeMapping(entry, attributes)
	return attrs, nil
}

func PullGroupInfo(group string, attr ...string) (models.GroupSimpleInfo, *customError.Error) {
	groupInfo := models.GroupSimpleInfo{}
	group = LDAP_STRING_REPLACE.Replace(group)

	/*
		"cn",
		"distinguishedName",
		"sAMAccountName",
		"description",
		"info",
	*/

	results, ldapError := SearchWithFilter(
		fmt.Sprintf("(&(objectCategory=group)(sAMAccountName=%s))", group),
		attr...,
	)

	if ldapError != nil {
		logger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return groupInfo, &cError
	}

	if results == nil || len(results.Entries) == 0 {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO GROUP FOUND FOR: %s", group))
		logger.Errorf("%v", cError)
		logger.Debug(results.Entries)
		return groupInfo, &cError
	}

	if len(results.Entries) > 1 {
		logger.Debugf("Multiple Group Entries: %+v\n Input: %s", results.Entries, group)
	}

	entry := results.Entries[0]

	return models.ToGroupSimpleInfo(entry), nil
}

func PullGroupInfoByDN(groupDN string, attr ...string) (models.GroupSimpleInfo, *customError.Error) {
	groupInfo := models.GroupSimpleInfo{}
	groupDN = LDAP_STRING_REPLACE.Replace(groupDN)

	/*
		"cn",
		"distinguishedName",
		"sAMAccountName",
		"description",
		"info",
	*/

	results, ldapError := SearchWithFilter(
		fmt.Sprintf("(&(objectCategory=group)(cn=%s))", groupDN),
		attr...,
	)

	if ldapError != nil {
		logger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return groupInfo, &cError
	}

	if results == nil || len(results.Entries) == 0 {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO GROUP FOUND FOR: %s", groupDN))
		logger.Errorf("%v", cError)
		logger.Debug(results.Entries)
		return groupInfo, &cError
	}

	if len(results.Entries) > 1 {
		logger.Debugf("Multiple Group Entries: %+v\n Input: %s", results.Entries, groupDN)
	}

	entry := results.Entries[0]

	return models.ToGroupSimpleInfo(entry), nil
}

func AddUsersToGroup(group string, newMembers []string) (map[string]models.GroupModifyResults, *customError.Error) {
	var results map[string]models.GroupModifyResults = make(map[string]models.GroupModifyResults)

	usersDN, cError := GetUsersDN(newMembers)
	if cError != nil {
		logger.Error(cError.ToString())
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

func RemoveUsersFromGroup(group string, members []string) (map[string]models.GroupModifyResults, *customError.Error) {
	var cError *customError.Error
	var results map[string]models.GroupModifyResults = make(map[string]models.GroupModifyResults)

	var usersDN map[string]string
	if usersDN, cError = GetUsersDN(members); cError != nil {
		logger.Error(cError.ToString())
		return results, cError
	}

	var groupDN string
	if groupDN, cError = GetGroupDN(group); cError != nil {
		logger.Error(cError.ToString())
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
				logger.Errorf("Could not find: %s - For a DN search", group)
				return groupsDN, cError
			}
			continue
		}

		groupsDN[group] = DN
	}

	return groupsDN, nil
}

func GetGroupDN(group string) (string, *customError.Error) {

	group = LDAP_STRING_REPLACE.Replace(group)

	results, ldapError := SearchWithFilter(fmt.Sprintf("(&(objectCategory=group)(|(sAMAccountName=%s)(cn=%s)))", group, group), DISTINGUISHED_NAME)

	if ldapError != nil {
		logger.Error(ldapError)
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
		logger.Error(cError.ToString())
		return cError
	}
	defer l.Close()
	defer l.Unbind()

	addRequest := ldap.NewModifyRequest(groupDN, nil)
	addRequest.Add("member", []string{user})
	ldapError := l.Modify(addRequest)

	if ldapError != nil {
		logger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return &cError
	}

	return nil
}

// * Sends a modify request to remove a member from group
func ModifyGroupRemoveMember(groupDN, user string) *customError.Error {
	l, cError := connectToLDAP()

	if cError != nil {
		logger.Error(cError.ToString())
		return cError
	}
	defer l.Close()
	defer l.Unbind()

	addRequest := ldap.NewModifyRequest(groupDN, nil)
	addRequest.Delete("member", []string{user})
	ldapError := l.Modify(addRequest)

	if ldapError != nil {
		logger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return &cError
	}

	return nil
}

// * Retrives all the members of a specific group
//
// NOTE: Total Retrive per request 1500
func GetAllMembers(group string) ([]string, *customError.Error) {
	start := 0
	end := 1499

	var final_memeber []string = []string{}
	var temp_members []string = []string{}

	l, cError := connectToLDAP()

	if cError != nil {
		logger.Error(cError.ToString())
		return final_memeber, cError
	}

	defer l.Close()
	defer l.Unbind()

	for {
		config := DefaultSearchConfig().SetFilter(fmt.Sprintf("(&(objectCategory=group)(cn=%s))", group)).SetAttributes([]string{memberRangeAtrribute(start, end)}).SetControl(nil)
		results, ldapError := config.Search()

		if ldapError != nil {
			logger.Error(ldapError)
			cError := customError.LDAP_ERROR.NewError(ldapError)
			return final_memeber, &cError
		}

		if results == nil {
			logger.Error("LDAP RETURN RESULTS WERE NIL")
			cError := customError.LDAP_ERROR.NewMessage("RESULTS WERE NIL FROM LDAP SEARCH")
			return final_memeber, &cError
		}

		if len(results.Entries) == 0 {
			logger.Errorf("GROUP NOT FOUND FOR: %s", group)
			cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO GROUPS FOUND FOR: %s", group))
			return final_memeber, &cError
		}

		result := results.Entries[0]

		for _, attr := range result.Attributes {
			temp_members = attr.Values
			final_memeber = append(final_memeber, temp_members...)
		}

		if len(temp_members) != 1500 {
			logger.Infof("END OF MEMBERS FOR GROUP: %s", group)
			break
		}

		if RATE_LIMIT < 0 {
			logger.Errorf("RATE LIMIT WAS REACH FOR MEMBERS: %s", group)
			cError := customError.NewError(http.StatusInternalServerError, "RATE_LIMIT_HIT", fmt.Sprintf("RATE LIMIT WAS REACH FOR MEMBERS: %s", group))
			return []string{}, &cError
		}

		start += 1500
		end += 1500
	}

	logger.Info("Total Members: ", len(final_memeber))

	return final_memeber, nil

}
