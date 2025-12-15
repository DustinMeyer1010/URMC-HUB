package ad

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/go-ldap/ldap/v3"
)

type LDAPSearchConfig struct {
	BaseDN    string
	Attribute []string
	Scope     int
	Deref     int
	SizeLimit int
	TimeLimit int
	TypesOnly bool
	Filter    string
	Control   []ldap.Control
}

// Creates a LDAPSearchConfig for performing a LDAP Search
func SearchConfig(filter string, attribute ...string) LDAPSearchConfig {

	return LDAPSearchConfig{
		BaseDN:    global.BASEDN,
		Scope:     ldap.ScopeWholeSubtree,
		Deref:     ldap.NeverDerefAliases,
		Attribute: attribute,
		Filter:    filter,
		SizeLimit: 0,
		TimeLimit: 0,
		TypesOnly: false,
		Control:   []ldap.Control{ldap.NewControlPaging(100)},
	}

}

// Searches ldap based on the config struct that is provided
func (config LDAPSearchConfig) Search(conn *ldap.Conn) (*ldap.SearchResult, error) {
	searchRequest := ldap.NewSearchRequest(
		config.BaseDN,
		config.Scope,
		config.Deref,
		config.SizeLimit,
		config.TimeLimit,
		config.TypesOnly,
		config.Filter,
		config.Attribute,
		config.Control,
	)

	return conn.Search(searchRequest)
}

func SearchAllByCategory(category, attribute, value string, attrs ...string) (*ldap.SearchResult, error) {

	conn, cError := connectToLDAP()

	if cError != nil {
		return nil, cError.GetErrorValue()
	}

	defer conn.Close()
	defer conn.Unbind()

	filter := fmt.Sprintf("(&(objectCategory=%s)(%s=%s*))", category, attribute, value)

	ldapConfig := SearchConfig(filter, attrs...)

	return ldapConfig.Search(conn)
}

func SearchByCategory(category, attribute, value string, attrs ...string) (*ldap.SearchResult, error) {

	conn, cError := connectToLDAP()

	if cError != nil {
		return nil, cError.GetErrorValue()
	}

	defer conn.Close()
	defer conn.Unbind()

	filter := fmt.Sprintf("(&(objectCategory=%s)(%s=%s))", category, attribute, value)

	ldapConfig := SearchConfig(filter, attrs...)

	return ldapConfig.Search(conn)
}

func SearchWithFilter(filter string, attrs ...string) (*ldap.SearchResult, error) {
	conn, cError := connectToLDAP()

	if cError != nil {
		return nil, cError.GetErrorValue()
	}

	defer conn.Close()
	defer conn.Unbind()

	ldapConfig := SearchConfig(filter)

	return ldapConfig.Search(conn)
}
