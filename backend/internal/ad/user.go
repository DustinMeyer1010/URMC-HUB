package ad

import (
	"fmt"
	"strings"
	"sync"

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

	if len(results.Entries) == 0 {
		return user, err
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
