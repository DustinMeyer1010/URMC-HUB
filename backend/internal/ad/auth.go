package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

func connectToLDAP() (*ldap.Conn, error) {

	if global.Username == "" || global.Password == "" {
		cError := errs.UNAUTHORIZED
		return nil, &cError
	}

	l, ldapError := ldap.DialURL(global.URMC_LDAP)

	if ldapError != nil {
		logger.Error(ldapError)
		cError := errs.LDAP_ERROR.NewError(ldapError)
		return nil, &cError
	}

	ldapError = l.Bind(formatUsername(global.Username), global.Password)

	if ldapError != nil {
		cError := errs.UNAUTHORIZED
		return nil, &cError
	}

	return l, nil
}

func Login(user models.UserLogin) error {

	l, ldapError := ldap.DialURL(global.URMC_LDAP)

	if ldapError != nil {
		logger.Error(ldapError)
		cError := errs.LDAP_ERROR.NewError(ldapError)
		return &cError
	}

	ldapError = l.Bind(formatUsername(user.Username), user.Password)

	if ldapError != nil {
		logger.Error(ldapError)
		cError := errs.LDAP_ERROR.NewError(ldapError)
		return &cError
	}

	global.Username = user.Username
	global.Password = user.Password

	CreatePresistantLdapConnection()
	LDAP_CONNECTION.StartHeartBeat()

	return nil

}

func CreatePresistantLdapConnection() error {
	conn, err := connectToLDAP()

	if err != nil {
		return err
	}

	LDAP_CONNECTION.Conn = conn

	return nil
}

func ConnectToServer(URL string) (*ldap.Conn, error) {
	if global.Username == "" || global.Password == "" {
		cError := errs.NOT_LOGGED_IN
		return nil, &cError
	}

	l, ldapError := ldap.DialURL(URL)

	if ldapError != nil {
		logger.Error(ldapError)
		cError := errs.LDAP_ERROR.NewError(ldapError)
		return nil, &cError
	}

	ldapError = l.Bind(fmt.Sprintf("%s\\%s", global.USERNAME_PREFIX, global.Username), global.Password)

	if ldapError != nil {
		logger.Error(ldapError)
		cError := errs.LDAP_ERROR.NewError(ldapError)
		return nil, &cError
	}

	return l, nil
}

func Verify() error {
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
