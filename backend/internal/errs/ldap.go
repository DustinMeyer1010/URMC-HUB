package errs

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

func LdapErrorToModifyResult(ldapError error) models.GroupModifyResults {

	if e, ok := ldapError.(*ldap.Error); ok {
		switch e.ResultCode {
		case ldap.LDAPResultUnwillingToPerform:
			return models.GroupModifyResults{}
		case ldap.LDAPResultInsufficientAccessRights:
			return models.GroupModifyResults{}
		default:
			return models.GroupModifyResults{}
		}
	} else {
		return models.GroupModifyResults{}
	}

}
