package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

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
