package ad

import (
	"fmt"
)

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
