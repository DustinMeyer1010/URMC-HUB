package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

// Deprecated: Will be moved over to lookupComputer
func PullComputerInformation(computer string) (models.ComputerSimpleInfo, error) {
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
		cError := errs.LDAP_ERROR.NewError(ldapError)
		return models.ComputerSimpleInfo{}, &cError
	}

	if results == nil || len(results.Entries) == 0 {
		cError := errs.NOT_FOUND.NewMessage(fmt.Sprintf("NO COMPUTER FOUND FOR: %s", computer))
		return models.ComputerSimpleInfo{}, &cError
	}

	return models.ToComputerSimpleInfo(results.Entries[0]), nil

}
