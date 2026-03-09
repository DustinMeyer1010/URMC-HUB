package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/go-ldap/ldap/v3"
)

// Looks up the user with the provided disgushiedName. Returns a map, where the key is
// the attrubutes provided & values being what active directory returned for that attrubtue.
// If attribute is not found it will just have an empty values with attribute being the key.
func LookupUser(userDN string, attributes ...string) (map[string][]string, error) {

	attrs := map[string][]string{}

	config := UserSearchConfig().
		SetAttributes(attributes).
		SetBaseDN(userDN)

	sr, ldapError := config.Search()

	if ldapError != nil {
		cError := errs.LDAP_ERROR.NewMessage(ldapError.Error())
		return map[string][]string{}, &cError
	}

	var entry *ldap.Entry

	if len(sr.Entries) > 0 {
		entry = sr.Entries[0]
	} else {
		return map[string][]string{}, &errs.NOT_FOUND
	}

	attrs = ExtractAttributes(entry, attributes)

	return attrs, nil

}

// Sanitizes a search string and executes a broad user query.
// It matches the input against standard identity attributes and populates the
// provided 'users' pointer with a collection of attribute maps defined by
// the 'attributes' parameter.
func SearchAllUserNew(users *[]map[string][]string, searchValue string, attributes ...string) error {

	searchValue = LDAP_STRING_REPLACE.Replace(searchValue)
	filter := fmt.Sprintf("(&(objectCategory=user)(|(anr=%s)(URID=%s)))", searchValue, searchValue)

	searchConfig := DefaultSearchConfig().
		SetFilter(filter).
		SetAttributes(attributes)

	searchResults, ldapError := searchConfig.Search()

	if cError := checkSearchErrors(ldapError, searchResults); cError != nil {
		return cError
	}

	*users = ExtractMultipleEntriesAtrributes(searchResults.Entries, attributes)

	return nil

}

// Creates a moadify request on the group and then adds the ldap account
// to that groups. Returns an error based on the results of the ldap
// modify request
func GroupAddToUser(userDN, groupDN string) error {

	modifyConfig := NewDefaultModifyConfig(groupDN)

	return modifyConfig.Add(userDN)

}

// Creates a modify request on the group and then adds the ldap account
// to that groups. Returns an error based on the results of the ldap
// modify request
func GroupRemoveFromUser(userDN, groupDN string) error {
	modifyConfig := NewDefaultModifyConfig(groupDN)

	return modifyConfig.Remove(userDN)
}
