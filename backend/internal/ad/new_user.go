package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

// Looks up the user with the provided disgushiedName. Returns a map, where the key is
// the attrubutes provided & values being what active directory returned for that attrubtue.
// If attribute is not found it will just have an empty values with attribute being the key.
func LookupUser(userDN string, attributes ...string) (map[string][]string, *customError.Error) {

	attrs := map[string][]string{}

	config := UserSearchConfig().
		SetAttributes(attributes).
		SetBaseDN(userDN)

	sr, ldapError := config.Search()

	if ldapError != nil {
		cError := customError.LDAP_ERROR.NewMessage(ldapError.Error())
		return map[string][]string{}, &cError
	}

	var entry *ldap.Entry

	if len(sr.Entries) > 0 {
		entry = sr.Entries[0]
	} else {
		return map[string][]string{}, &customError.NOT_FOUND
	}

	attrs = createAttributeMapping(entry, attributes)

	return attrs, nil

}

// Looks up users that match the search values. This does a search based on the filters
// anr & URID. Returns a map with the key being the attribute provided and the value being
// what active directory returns. If no attribute found for active directory then the value
// will be an empty string.
func SearchAllUsers(searchValue string, attr ...string) ([]models.UserSimpleInfo, *customError.Error) {

	matches := make([]models.UserSimpleInfo, 0)
	searchValue = LDAP_STRING_REPLACE.Replace(searchValue)
	filter := fmt.Sprintf("(&(objectCategory=user)(|(anr=%s)(URID=%s)))", searchValue, searchValue)

	ldapConfig := DefaultSearchConfig().SetFilter(filter).SetAttributes(attr)

	results, ldapError := ldapConfig.Search()

	if cError := checkSearchErrors(ldapError, results); cError != nil {
		return matches, cError
	}

	for _, entry := range results.Entries {
		var user models.UserSimpleInfo
		user.FillAttributes(entry)
		matches = append(matches, user)
	}

	return matches, nil
}

func SearchAllUserNew(users *[]map[string][]string, searchValue string, attributes ...string) *customError.Error {

	searchValue = LDAP_STRING_REPLACE.Replace(searchValue)
	filter := fmt.Sprintf("(&(objectCategory=user)(|(anr=%s)(URID=%s)))", searchValue, searchValue)

	searchConfig := DefaultSearchConfig().
		SetFilter(filter).
		SetAttributes(attributes)

	searchResults, ldapError := searchConfig.Search()

	if cError := checkSearchErrors(ldapError, searchResults); cError != nil {
		return cError
	}

	// REFACTOR: Create all the mapping for each attribute
	for _, e := range searchResults.Entries {
		attr := make(map[string][]string)
		for i, a := range convertAliases(attributes) {
			attr[attributes[i]] = e.GetAttributeValues(a)
		}
		*users = append(*users, attr)
	}

	return nil

}

func GroupAddToUser(userDN, groupDN string) *customError.Error {

	modifyConfig := NewDefaultModifyConfig(groupDN)

	return modifyConfig.Add(userDN)

}

func GroupRemoveFromUser(userDN, groupDN string) *customError.Error {
	modifyConfig := NewDefaultModifyConfig(groupDN)

	return modifyConfig.Remove(userDN)
}
