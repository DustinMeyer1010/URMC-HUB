package ad

import (
	"fmt"
	"strings"
	"sync"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

func SearchAllUsers(searchValue string) ([]models.UserSimpleInfo, *customError.Error) {

	matches := make([]models.UserSimpleInfo, 0)

	l, cError := connectToLDAP()

	if cError != nil {
		return matches, cError
	}

	defer l.Close()
	defer l.Unbind()

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

	results, ldapError := ldapConfig.Search(l)

	if ldapError != nil {
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return matches, &cError
	}

	if results == nil {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO USERS FOUND FOR: %s", searchValue))
		return matches, &cError
	}

	for _, entry := range results.Entries {
		var user models.UserSimpleInfo
		user.FillAttributes(entry)
		matches = append(matches, user)
	}

	return matches, nil
}

func PullUserInformation(searchValue string) (models.UserFullInfo, *customError.Error) {

	var user models.UserFullInfo

	catagory := "user"
	attribute := "SAMAccountName"

	results, ldapError := SearchAllByCategory(
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

	if ldapError != nil {
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return user, &cError
	}

	if results == nil {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO USER FOUND FOR: %s", searchValue))
		return user, &cError
	}

	foundUser := results.Entries[0]

	groups := user.FillAttributes(foundUser)
	memberOfChan := make(chan models.GroupSimpleInfo, len(groups))
	var wg sync.WaitGroup
	for _, group := range groups {
		wg.Add(1)
		go func(group string) {
			defer wg.Done()
			g := strings.Split(group, ",")[0][3:]
			found, _ := PullGroupInfo(g)
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

	user.MemberOf = memberOf

	return user, nil
}

// TODO: Update to only take one user as the argument
func AddGroup(users []string, groups []string) ([]models.GroupModifyResults, *customError.Error) {

	l, cError := connectToLDAP()
	if cError != nil {
		return []models.GroupModifyResults{}, cError
	}
	defer l.Close()
	defer l.Unbind()

	var usersDN map[string]string
	if usersDN, cError = GetUsersDN(users); cError != nil {
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

func RemoveGroup(users []string, groups []string) ([]models.GroupModifyResults, *customError.Error) {

	l, cError := connectToLDAP()
	if cError != nil {
		return []models.GroupModifyResults{}, cError
	}
	defer l.Close()
	defer l.Unbind()

	usersDN, cError := GetUsersDN(users)

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
		results, ldapError := SearchByCategory("user", "SAMaccountName", user, "dn")

		if ldapError != nil {
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

func UserDetails(input string) ([]models.UserDetails, error) {

	foundUser := make([]models.UserDetails, 0)

	config := SearchConfig(
		fmt.Sprintf("(&(objectClass=user)(anr=%s))", input),
		"name",
		"mail",
		"sAMAccountName",
	)

	if persistConn == nil {
		return []models.UserDetails{}, fmt.Errorf("connection to Ldap is nil")
	}

	results, err := config.Search(persistConn)

	if err != nil {
		return []models.UserDetails{}, err
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
