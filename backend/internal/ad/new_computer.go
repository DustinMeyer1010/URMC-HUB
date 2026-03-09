package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func SearchAllComputers(searchValue string, attributes ...string) ([]models.ComputerSimpleInfo, error) {
	computers := []models.ComputerSimpleInfo{}
	searchValue = LDAP_STRING_REPLACE.Replace(searchValue)

	searchConfig := DefaultSearchConfig().
		SetFilter(fmt.Sprintf("(&(objectClass=computer)(anr=%s*))", searchValue)).
		SetAttributes(attributes)

	searchResults, ldapError := searchConfig.Search()

	if searchResults == nil {
		cError := &errs.NOT_FOUND
		return computers, cError
	}

	if ldapError != nil {
		logger.Error(ldapError)
		cError := errs.LDAP_ERROR.NewError(ldapError)
		return computers, &cError
	}

	for _, entry := range searchResults.Entries {
		computers = append(computers, models.ToComputerSimpleInfo(entry))
	}

	return computers, nil
}

func SearchAllComputersNew(computers *[]map[string][]string, searchValue string, attributes ...string) error {
	searchValue = LDAP_STRING_REPLACE.Replace(searchValue)
	filter := fmt.Sprintf("(&(objectCategory=computer)(anr=%s))", searchValue)

	searchConfig := DefaultSearchConfig().
		SetFilter(filter).
		SetAttributes(attributes)

	searchResults, ldapError := searchConfig.Search()

	if cError := checkSearchErrors(ldapError, searchResults); cError != nil {
		return cError
	}

	*computers = ExtractMultipleEntriesAtrributes(searchResults.Entries, attributes)

	return nil
}

// Performs an LDAP search for a specific computer by its Distinguished Name.
// It returns a mapped collection of the requested attributes or a custom error if the
// computer is not found or the search fails.
func LookupComputer(computerDN string, attributes ...string) (map[string][]string, error) {

	searchConfig := ComputerSearchConfig().
		SetBaseDN(computerDN).
		SetAttributes(attributes)

	searchResults, ldapError := searchConfig.Search()

	if ldapError != nil {
		logger.Error(ldapError)
		cError := errs.LDAP_ERROR.NewError(ldapError)
		return map[string][]string{}, &cError
	}

	if searchResults == nil || len(searchResults.Entries) == 0 {
		cError := errs.NOT_FOUND.NewMessage(fmt.Sprintf("NO COMPUTER FOUND FOR: %s", computerDN))
		return map[string][]string{}, &cError
	}

	entry := searchResults.Entries[0]

	attrs := ExtractAttributes(entry, attributes)

	return attrs, nil
}
