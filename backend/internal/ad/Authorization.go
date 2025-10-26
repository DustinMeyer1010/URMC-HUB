package ad

import (
	"fmt"
	"log"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

func connectToLDAP() (l *ldap.Conn, err error) {

	l, err = ldap.DialURL(global.URMC_LDAP)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = l.Bind(formatUsername(global.Username), global.Password)

	return
}

func Login(user models.UserLogin) error {

	l, err := ldap.DialURL(global.URMC_LDAP)

	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Failed to connect to URMC ldap server")
	}

	err = l.Bind(formatUsername(user.Username), user.Password)

	if err != nil {
		return fmt.Errorf("invalid username or password")
	}

	global.Username = user.Username
	global.Password = user.Password

	return nil

}

func ConnectToServer(URL string) (*ldap.Conn, error) {
	l, err := ldap.DialURL(URL)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = l.Bind(fmt.Sprintf("%s\\%s", global.USERNAME_PREFIX, global.Username), global.Password)

	return l, err
}

func Verify() error {

	if global.Username == "" || global.Password == "" {
		return fmt.Errorf("no credentials provided")
	}

	_, err := connectToLDAP()

	if err != nil {
		return fmt.Errorf("invalid username or password")
	}

	return nil
}

func formatUsername(username string) string {
	return fmt.Sprintf(
		"%s\\%s",
		global.USERNAME_PREFIX,
		username,
	)
}
