package ad

import (
	"fmt"
	"sync"
	"time"

	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

var LDAP_CONNECTION LDAPConnection

type LDAPConnection struct {
	Conn  *ldap.Conn
	mutex sync.Mutex
}

// Maintains the connection with application and the ldap server
func (c *LDAPConnection) StartHeartBeat() {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			c.mutex.Lock()
			fmt.Println("Checking Connection ...")
			_, err := c.Conn.Search(ldap.NewSearchRequest("", ldap.ScopeBaseObject, ldap.NeverDerefAliases, 0, 0, false, "(objectClass=*)", []string{}, nil))
			if err != nil {
				fmt.Println("Restoring Connection")
				err := c.Restore()

				if err != nil {
					panic("something wrong with ldap")
				}
			}
			c.mutex.Unlock()
		}
	}()
}

// Restores the connection to the ldap server if it is lost
func (c *LDAPConnection) Restore() error {
	if c.Conn != nil {
		c.Conn.Close()
	}

	var newConn *ldap.Conn
	var err error

	for range 5 {
		newConn, err = connectToLDAP()
		if err == nil {
			break
		}
	}

	c.Conn = newConn

	return nil
}

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
