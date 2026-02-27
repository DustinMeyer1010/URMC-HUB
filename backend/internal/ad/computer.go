package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func SearchAllComputers(searchValue string, attrs ...attribute) ([]models.ComputerSimpleInfo, *customError.Error) {
	matches := []models.ComputerSimpleInfo{}
	searchValue = LDAP_STRING_REPLACE.Replace(searchValue)
	results, ldapError := SearchAllByCategory(
		"computer",
		"anr",
		searchValue,
		attrs...,
	)

	if results == nil {
		cError := &customError.NOT_FOUND
		return matches, cError
	}

	if ldapError != nil {
		logger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return matches, &cError
	}

	for _, entry := range results.Entries {
		matches = append(matches, models.ToComputerSimpleInfo(entry))
	}

	return matches, nil
}

func PullComputerInformation(computer string, attr ...attribute) (models.ComputerSimpleInfo, *customError.Error) {
	computer = LDAP_STRING_REPLACE.Replace(computer)
	results, ldapError := SearchByCategory(
		"computer",
		"name",
		computer,
		attr...,
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
