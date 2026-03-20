package ad

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/go-ldap/ldap/v3"
)

func checkSearchErrors(ldapErr error, results *ldap.SearchResult) error {

	if ldapErr != nil {
		logger.Error(ldapErr)
		cError := errs.LDAP_ERROR.NewError(ldapErr)
		return &cError
	}

	if results == nil || len(results.Entries) == 0 {
		cError := errs.NOT_FOUND.NewMessage("NOT FOUND")
		return &cError
	}

	return nil
}
