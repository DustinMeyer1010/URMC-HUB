package ad

import (
	"fmt"
	"log"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

func connectToLDAP() (l *ldap.Conn, err error) {

	l, err = ldap.DialURL("ldap://URMC-sh.rochester.edu/")

	if err != nil {
		log.Fatal(err)
		return
	}

	err = l.Bind(fmt.Sprintf("URMC-sh\\%s", global.Username), global.Password)

	return
}

func Login(user models.UserLogin) error {

	l, err := ldap.DialURL("ldap://URMC-sh.rochester.edu/")

	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Failed to Dail the Server")
	}

	urmcsh := fmt.Sprintf("URMC-sh\\%s", user.Username)

	err = l.Bind(urmcsh, user.Password)

	if err != nil {
		return fmt.Errorf("invalid Username or Password")
	}

	global.Username = user.Username
	global.Password = user.Password

	return nil

}

func Verify() error {

	if global.Username == "" || global.Password == "" {
		return fmt.Errorf("no Username or Password")
	}

	_, err := connectToLDAP()

	if err != nil {
		return fmt.Errorf("invalid username or password")
	}

	return nil
}
