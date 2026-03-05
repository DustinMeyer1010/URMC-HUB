package ad

import (
	"fmt"
	"strings"
	"sync"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

// Deprecated: To be replace with LookupUser
func PullUserInformation(searchValue string, attr ...string) (models.UserFullInfo, *customError.Error) {

	var user models.UserFullInfo

	catagory := "user"
	attribute := "SAMAccountName"

	searchValue = LDAP_STRING_REPLACE.Replace(searchValue)

	/*
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
		"URRoleStatus",
		"sn",
	*/

	results, ldapError := SearchAllByCategory(
		catagory,
		attribute,
		searchValue,
		attr...,
	)

	if ldapError != nil {
		logger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return user, &cError
	}

	if results == nil || len(results.Entries) == 0 {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO USER FOUND FOR: %s", searchValue))
		return user, &cError
	}

	foundUser := results.Entries[0]

	user.FillAttributes(foundUser)
	return user, nil
}

// Deprecated: Replaced by GetUserGroups
func PullUserMembersOf(username string) ([]models.GroupSimpleInfo, *customError.Error) {
	catagory := "user"
	attribute := "SAMAccountName"

	results, ldapError := SearchByCategory(
		catagory,
		attribute,
		username,
		MEMBER_OF,
	)

	if ldapError != nil {
		logger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return []models.GroupSimpleInfo{}, &cError
	}

	if results == nil || len(results.Entries) == 0 {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO USER FOUND FOR: %s", username))
		return []models.GroupSimpleInfo{}, &cError
	}

	groups := results.Entries[0].GetAttributeValues("memberOf")

	memberOfChan := make(chan models.GroupSimpleInfo, len(groups))
	var wg sync.WaitGroup
	for _, group := range groups {
		wg.Add(1)
		go func(group string) {
			defer wg.Done()
			g := strings.Split(group, ",")[0][3:]
			found, _ := PullGroupInfoByDN(g)
			memberOfChan <- found
		}(group)
	}

	go func() {
		wg.Wait()
		close(memberOfChan)
	}()

	// Collect all results into a slice
	var memberOf []models.GroupSimpleInfo
	for result := range memberOfChan {
		memberOf = append(memberOf, result)
	}

	return memberOf, nil

}

func AddGroup(username string, groups []string) ([]models.GroupModifyResults, *customError.Error) {

	l, cError := connectToLDAP()
	if cError != nil {
		return []models.GroupModifyResults{}, cError
	}
	defer l.Close()
	defer l.Unbind()

	var usersDN map[string]string
	if usersDN, cError = GetUsersDN([]string{username}); cError != nil {
		return []models.GroupModifyResults{}, cError
	}

	if len(usersDN) == 0 {
		cError := customError.NOT_FOUND.NewMessage("NO USER FOUND TO BECOME MEMEBER OF GROUP")
		return []models.GroupModifyResults{}, &cError
	}

	var groupsDN map[string]string
	if groupsDN, cError = GetGroupsDN(groups); cError != nil {
		return []models.GroupModifyResults{}, cError
	}

	var response []models.GroupModifyResults

	for group, groupDN := range groupsDN {
		for _, userDN := range usersDN {
			groupResult := models.GroupModifyResults{
				Group:      group,
				Successful: true,
				Message:    "Group Added",
			}
			cError = ModifyGroupNewMember(groupDN, userDN)
			if cError != nil {
				groupResult.Successful = false
				groupResult.Message = cError.Msg
				response = append(response, groupResult)
				continue
			}
			response = append(response, groupResult)
		}

	}

	return response, nil
}

func RemoveGroup(username string, groups []string) ([]models.GroupModifyResults, *customError.Error) {

	l, cError := connectToLDAP()
	if cError != nil {
		return []models.GroupModifyResults{}, cError
	}
	defer l.Close()
	defer l.Unbind()

	usersDN, cError := GetUsersDN([]string{username})

	if cError != nil {
		return []models.GroupModifyResults{}, cError
	}

	if len(usersDN) == 0 {
		cError := customError.NOT_FOUND.NewMessage("NO USER FOUND REMOVE FROM GROUP")
		return []models.GroupModifyResults{}, &cError
	}

	groupsDN, cError := GetGroupsDN(groups)

	if cError != nil {
		return []models.GroupModifyResults{}, cError
	}

	var response []models.GroupModifyResults

	// Create delete request for each group
	for group, groupDN := range groupsDN {
		for _, userDN := range usersDN {
			groupResult := models.GroupModifyResults{
				Group:      group,
				Successful: true,
				Message:    "Groups Removed",
			}
			cError := ModifyGroupRemoveMember(groupDN, userDN)
			if cError != nil {
				groupResult.Successful = false
				groupResult.Message += cError.Msg
				response = append(response, groupResult)
				continue
			}
			response = append(response, groupResult)
		}

	}

	return response, nil

}

func GetUsersDN(users []string) (map[string]string, *customError.Error) {
	var userDN map[string]string = make(map[string]string)
	l, cError := connectToLDAP()
	if cError != nil {
		return userDN, cError
	}
	defer l.Close()
	defer l.Unbind()

	for _, user := range users {
		results, ldapError := SearchByCategory("user", "SAMaccountName", user, DISTINGUISHED_NAME)

		if ldapError != nil {
			logger.Error(ldapError)
			cError := customError.LDAP_ERROR.NewError(ldapError)
			return userDN, &cError
		}

		if len(results.Entries) == 0 {
			continue
		}

		userDN[user] = results.Entries[0].DN
	}

	return userDN, nil
}

var persistConn *ldap.Conn

func UserDetails(input string, attr ...string) ([]models.UserDetails, error) {

	foundUser := make([]models.UserDetails, 0)

	/*
		"name",
		"mail",
		"sAMAccountName",
		"department",
	*/

	config := DefaultSearchConfig().SetFilter(fmt.Sprintf("(&(objectClass=user)(anr=%s))", input)).SetAttributes(attr)

	if persistConn == nil {
		return []models.UserDetails{}, fmt.Errorf("connection to Ldap is nil")
	}

	results, ldapError := config.Search()

	if ldapError != nil {
		logger.Error(ldapError)
		return []models.UserDetails{}, ldapError
	}

	if len(results.Entries) == 0 {
		return []models.UserDetails{}, fmt.Errorf("NOT_FOUND")
	}

	for _, entry := range results.Entries {
		var user models.UserDetails
		if entry.GetAttributeValue("mail") == "" {
			continue
		}
		user.FillAttributes(entry)
		foundUser = append(foundUser, user)
	}

	return foundUser, nil
}

func CreatePresistantConn() (err error) {
	var cError *customError.Error
	persistConn, cError = connectToLDAP()
	if cError != nil {
		return fmt.Errorf("%s", cError.Msg)
	}

	return nil
}

func DisconnectPresistantConn() error {
	err := persistConn.Unbind()

	if err != nil {
		return err
	}

	err = persistConn.Close()

	return err

}

type BulkSearchResult struct {
	Value       string
	UserDetails models.UserDetails
}

func BulkLookup(values []string) (unique []BulkSearchResult, duplicates [][]BulkSearchResult, notFound []BulkSearchResult, err error) {
	err = CreatePresistantConn()

	if err != nil {
		return
	}

	for _, value := range values {
		currentSearch := BulkSearchResult{Value: value}
		results, err := UserDetails(value)

		if err != nil && err.Error() != "NOT_FOUND" {

			notFound = append(notFound, currentSearch)
			continue
		}

		if len(results) == 0 {
			notFound = append(notFound, currentSearch)
			continue
		}

		if len(results) == 1 {
			currentSearch.UserDetails = results[0]
			unique = append(unique, currentSearch)
			continue
		}
		var dup []BulkSearchResult = make([]BulkSearchResult, 0)

		for _, result := range results {
			currentSearch = BulkSearchResult{Value: value, UserDetails: result}
			dup = append(dup, currentSearch)
		}

		duplicates = append(duplicates, dup)
	}

	DisconnectPresistantConn()

	return

}
