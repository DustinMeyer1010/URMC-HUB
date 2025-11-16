package ad

import (
	"fmt"
	"strings"
	"sync"

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

	if len(results.Entries) == 0 {
		return user, fmt.Errorf("%s", "No accounts found")
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

	return user, err
}

func AddGroup(users []string, groups []string) ([]models.GroupModifyResults, error) {

	l, err := connectToLDAP()
	if err != nil {
		return []models.GroupModifyResults{}, err
	}
	defer l.Close()
	defer l.Unbind()

	usersDN, err := GetUserDN(users, l)

	if err != nil {
		return []models.GroupModifyResults{}, err
	}

	groupsDN, err := GetGroupsDN(groups, l)

	if err != nil {
		return []models.GroupModifyResults{}, err
	}

	var response []models.GroupModifyResults

	for _, group := range groupsDN {
		groupResult := new(models.GroupModifyResults)
		groupResult.Group = group
		groupResult.Successful = true
		groupResult.Message = "All changes completed"
		addRequest := ldap.NewModifyRequest(group, nil)
		addRequest.Add("member", usersDN)
		groupAddError := l.Modify(addRequest)
		if groupAddError != nil {
			fmt.Printf("Failed to add user to %s\n", group)
			groupResult.Successful = false
			groupResult.Message = groupAddError.Error()
			response = append(response, *groupResult)
			continue
		}
		response = append(response, *groupResult)
	}

	return response, err
}

func RemoveGroup(users []string, groups []string) ([]models.GroupModifyResults, error) {

	l, err := connectToLDAP()
	if err != nil {
		return []models.GroupModifyResults{}, err
	}
	defer l.Close()
	defer l.Unbind()

	usersDN, err := GetUserDN(users, l)

	if err != nil {
		return []models.GroupModifyResults{}, err
	}

	groupsDN, err := GetGroupsDN(groups, l)

	if err != nil {
		return []models.GroupModifyResults{}, err
	}

	var response []models.GroupModifyResults

	// Create delete request for each group
	for _, group := range groupsDN {

		groupResult := new(models.GroupModifyResults)
		groupResult.Group = group
		groupResult.Successful = true
		groupResult.Message = "All changes completed"
		deleteRequest := ldap.NewModifyRequest(group, nil)
		deleteRequest.Delete("member", usersDN)
		groupRemoveErr := l.Modify(deleteRequest)
		if groupRemoveErr != nil {
			fmt.Println("Failed to remove user from " + group)
			groupResult.Successful = false
			groupResult.Message += groupRemoveErr.Error()
			response = append(response, *groupResult)
			continue
		}
		response = append(response, *groupResult)
	}

	return response, err

}

func GetUserDN(users []string, l *ldap.Conn) ([]string, error) {

	var userDN []string

	for _, user := range users {
		config := SearchConfig(
			fmt.Sprintf("(&(objectClass=user)(SAMaccountName=%s))", user),
			"dn",
		)

		results, err := config.Search(l)

		if err != nil || len(results.Entries) == 0 {
			break
		}

		userDN = append(userDN, results.Entries[0].DN)
	}

	return userDN, nil
}

func GetGroupsDN(groups []string, l *ldap.Conn) ([]string, error) {

	var groupsDN []string

	for _, group := range groups {
		config := SearchConfig(
			fmt.Sprintf("(&(objectClass=group)(cn=%s))", group),
			"dn",
		)

		results, err := config.Search(l)

		if err != nil || len(results.Entries) == 0 {
			break
		}
		groupsDN = append(groupsDN, results.Entries[0].DN)
	}

	return groupsDN, nil
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
	persistConn, err = connectToLDAP()
	if err != nil {
		return err
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
