package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

func connectToLDAP() (*ldap.Conn, *customError.Error) {

	if global.Username == "" || global.Password == "" {
		cError := customError.UNAUTHORIZED
		return nil, &cError
	}

	l, ldapError := ldap.DialURL(global.URMC_LDAP)

	if ldapError != nil {
		logger.ServerLogger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return nil, &cError
	}

	ldapError = l.Bind(formatUsername(global.Username), global.Password)

	if ldapError != nil {
		cError := customError.UNAUTHORIZED
		return nil, &cError
	}

	return l, nil
}

func Login(user models.UserLogin) *customError.Error {

	l, ldapError := ldap.DialURL(global.URMC_LDAP)

	if ldapError != nil {
		logger.ServerLogger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return &cError
	}

	ldapError = l.Bind(formatUsername(user.Username), user.Password)

	if ldapError != nil {
		logger.ServerLogger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return &cError
	}

	global.Username = user.Username
	global.Password = user.Password

	return nil

}

func ConnectToServer(URL string) (*ldap.Conn, *customError.Error) {
	if global.Username == "" || global.Password == "" {
		cError := customError.NOT_LOGGED_IN
		return nil, &cError
	}

	l, ldapError := ldap.DialURL(URL)

	if ldapError != nil {
		logger.ServerLogger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return nil, &cError
	}

	ldapError = l.Bind(fmt.Sprintf("%s\\%s", global.USERNAME_PREFIX, global.Username), global.Password)

	if ldapError != nil {
		logger.ServerLogger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return nil, &cError
	}

	return l, nil
}

func Verify() *customError.Error {
	_, cError := connectToLDAP()
	return cError
}

func formatUsername(username string) string {
	return fmt.Sprintf(
		"%s\\%s",
		global.USERNAME_PREFIX,
		username,
	)
}
