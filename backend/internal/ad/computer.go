package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func SearchAllComputers(searchValue string, attributes ...string) ([]models.ComputerSimpleInfo, *customError.Error) {
	computers := []models.ComputerSimpleInfo{}
	searchValue = LDAP_STRING_REPLACE.Replace(searchValue)

	searchConfig := DefaultSearchConfig().
		SetFilter(fmt.Sprintf("(&(objectClass=computer)(anr=%s*))", searchValue)).
		SetAttributes(attributes)

	searchResults, ldapError := searchConfig.Search()

	if searchResults == nil {
		cError := &customError.NOT_FOUND
		return computers, cError
	}

	if ldapError != nil {
		logger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return computers, &cError
	}

	for _, entry := range searchResults.Entries {
		computers = append(computers, models.ToComputerSimpleInfo(entry))
	}

	return computers, nil
}

func LookupComputer(computerDN string, attributes ...string) (map[string][]string, *customError.Error) {

	searchConfig := ComputerSearchConfig().
		SetBaseDN(computerDN).
		SetAttributes(attributes)

	searchResults, ldapError := searchConfig.Search()

	if ldapError != nil {
		logger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return map[string][]string{}, &cError
	}

	if searchResults == nil || len(searchResults.Entries) == 0 {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO COMPUTER FOUND FOR: %s", computerDN))
		return map[string][]string{}, &cError
	}

	entry := searchResults.Entries[0]

	attrs := createAttributeMapping(entry, attributes)

	return attrs, nil

}

func PullComputerInformation(computer string) (models.ComputerSimpleInfo, *customError.Error) {
	computer = LDAP_STRING_REPLACE.Replace(computer)
	results, ldapError := SearchByCategory(
		"computer",
		"name",
		computer,
		"name",
		"operatingSystem",
		"distinguishedName",
	)

	if ldapError != nil {
		logger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return models.ComputerSimpleInfo{}, &cError
	}

	if results == nil || len(results.Entries) == 0 {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO COMPUTER FOUND FOR: %s", computer))
		return models.ComputerSimpleInfo{}, &cError
	}

	return models.ToComputerSimpleInfo(results.Entries[0]), nil

}
