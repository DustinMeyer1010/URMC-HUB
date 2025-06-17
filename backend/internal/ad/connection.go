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

	err = l.Bind(fmt.Sprintf("URMC-sh\\%s", user.Username), user.Password)

	if err != nil {
		return fmt.Errorf("Invalid Username or Password")
	}

	return nil

}
