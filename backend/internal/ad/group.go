package ad

import (
	"fmt"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

var LDAP_STRING_REPLACE = strings.NewReplacer(
	`\`, `\5c`,
	`*`, `\2a`,
	`(`, `\28`,
	`)`, `\29`,
	"\x00", `\00`,
)

var RATE_LIMIT = 100

// Sanitizes a search string and executes a broad group query.
// It matches the input against standard identity attributes and populates the
// provided 'groups' pointer with a collection of attribute maps defined by
// the 'attributes' parameter.
func SearchAllGroupsNew(groups *[]map[string][]string, searchValue string, attributes ...string) error {
	searchValue = LDAP_STRING_REPLACE.Replace(searchValue)
	filter := fmt.Sprintf("(&(objectCategory=group)(anr=%s))", searchValue)

	searchConfig := DefaultSearchConfig().
		SetFilter(filter).
		SetAttributes(attributes)

	searchResults, ldapError := searchConfig.Search()

	if cError := checkSearchErrors(ldapError, searchResults); cError != nil {
		return cError
	}

	*groups = ExtractMultipleEntriesAtrributes(searchResults.Entries, attributes)

	return nil

}

// Performs an LDAP search for a specific group by its Distinguished Name.
// It returns a mapped collection of the requested attributes or a custom error if the
// group is not found or the search fails.
func LookupGroup(groupDN string, attributes ...string) (map[string][]string, error) {

	searchConfig := GroupSearchConfig().
		SetBaseDN(groupDN).
		SetAttributes(attributes)

	searchResults, ldapError := searchConfig.Search()

	if cError := checkSearchErrors(ldapError, searchResults); cError != nil {
		return map[string][]string{}, cError
	}

	entry := searchResults.Entries[0]

	attrs := ExtractAttributes(entry, attributes)
	return attrs, nil
}

func GetGroupMembers(groupDN string, start, end int) (map[string][]string, error) {
	attr := make(map[string][]string, 0)
	memberAttr := []string{fmt.Sprintf("member;range=%d-%d", start, end)}

	searchConfig := DefaultSearchConfig().SetBaseDN(groupDN).SetAttributes(memberAttr)

	searchResults, ldapError := searchConfig.Search()

	if cError := checkSearchErrors(ldapError, searchResults); cError != nil {
		return map[string][]string{}, cError
	}

	entry := searchResults.Entries[0]

	for _, a := range entry.Attributes {
		attr["members"] = a.Values
	}

	return attr, nil

}

func UsersRemoveFromGroup(usersDN []string, groupDN string) {
	modifyConfig := NewDefaultModifyConfig(groupDN)

	for _, u := range usersDN {
		modifyConfig.Add(u)
	}
}

func UsersAddToGroup(usersDN []string, groupDN string) []models.GroupModifyResults {

	results := make([]models.GroupModifyResults, len(usersDN))

	modifyConfig := NewDefaultModifyConfig(groupDN)

	for _, u := range usersDN {
		searchConfig := UserSearchConfig().SetBaseDN(u).SetAttributes([]string{"samAccountName"})
		searchResult, ldapError := searchConfig.EntryExist()

		if ldapError != nil {
			results = append(results, models.GroupModifyResults{
				Group:      groupDN,
				User:       "NOT_FOUND",
				Message:    fmt.Sprintf("Failed to find user with DN: %s", u),
				Successful: false},
			)
			continue
		}

		ldapError = modifyConfig.Remove(u)

		if ldapError != nil {
			results = append(results, models.GroupModifyResults{
				Group:      groupDN,
				User:       u,
				Message:    fmt.Sprintf("Failed to find user with DN: %s", u),
				Successful: false},
			)
			continue
		}

		results = append(results, models.GroupModifyResults{Group: groupDN, User: searchResult.GetAttributeValue("samAccountName"), Message: "Successful Add", Successful: false})
	}

	return results

}
