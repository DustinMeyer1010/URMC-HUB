package ad

import (
	"fmt"
	"sync"
	"time"

	"github.com/gen2brain/dlgs"
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

func (c *LDAPConnection) CheckConnection() {
	_, err := c.Conn.Search(ldap.NewSearchRequest("", ldap.ScopeBaseObject, ldap.NeverDerefAliases, 0, 0, false, "(objectClass=*)", []string{}, nil))
	if err != nil {
		c.Restore()
	}
}

// Restores the connection to the ldap server if it is lost
func (c *LDAPConnection) Restore() error {
	if c.Conn != nil {
		c.Conn.Close()
	}
	var connectionEstablished bool = false
	var newConn *ldap.Conn
	var err error

	for range 5 {
		newConn, err = connectToLDAP()
		if err == nil {
			connectionEstablished = true
			break
		}
	}

	if !connectionEstablished {
		dlgs.Error("Server Connection Failed", "Unable to reach URMC.rochester.edu ldap servers. Please make sure you are connected to URMC network or GlobalProtect.")
	}

	c.Conn = newConn

	return nil
}
