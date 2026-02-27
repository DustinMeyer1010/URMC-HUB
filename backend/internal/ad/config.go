package ad

import (
	"bufio"
	"fmt"
	"os"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/go-ldap/ldap/v3"
)

type LDAPSearchConfig struct {
	BaseDN     string
	Attributes []string
	Scope      int
	Deref      int
	SizeLimit  int
	TimeLimit  int
	TypesOnly  bool
	Filter     string
	Control    []ldap.Control
}

// Creates a LDAPSearchConfig for performing a LDAP Search
func SearchConfig(filter string, attrs ...attribute) LDAPSearchConfig {

	return LDAPSearchConfig{
		BaseDN:     global.BASEDN,
		Scope:      ldap.ScopeWholeSubtree,
		Deref:      ldap.NeverDerefAliases,
		Attributes: attributesToStrings(attrs...),
		Filter:     filter,
		SizeLimit:  0,
		TimeLimit:  0,
		TypesOnly:  false,
		Control:    []ldap.Control{ldap.NewControlPaging(100)},
	}

}

// Searches ldap based on the config struct that is provided
func (config LDAPSearchConfig) Search() (*ldap.SearchResult, error) {

	conn, cError := connectToLDAP()
	if cError != nil {
		return nil, cError.GetErrorValue()
	}

	defer conn.Close()
	defer conn.Unbind()

	searchRequest := ldap.NewSearchRequest(
		config.BaseDN,
		config.Scope,
		config.Deref,
		config.SizeLimit,
		config.TimeLimit,
		config.TypesOnly,
		config.Filter,
		config.Attributes,
		config.Control,
	)

	return conn.Search(searchRequest)
}

func SearchAllByCategory(category, attribute, value string, attrs ...attribute) (*ldap.SearchResult, error) {

	filter := fmt.Sprintf("(&(objectCategory=%s)(%s=%s*))", category, attribute, value)

	ldapConfig := SearchConfig(filter, attrs...)

	return ldapConfig.Search()
}

func SearchByCategory(category, attribute, value string, attrs ...attribute) (*ldap.SearchResult, error) {

	filter := fmt.Sprintf("(&(objectCategory=%s)(%s=%s))", category, attribute, value)

	ldapConfig := SearchConfig(filter, attrs...)

	return ldapConfig.Search()
}

func SearchWithFilter(filter string, attrs ...attribute) (*ldap.SearchResult, error) {

	ldapConfig := SearchConfig(filter, attrs...)

	return ldapConfig.Search()
}

func SearchByUserDN(UserDN string, attrs ...attribute) (*ldap.SearchResult, error) {

	ldapConfig := SearchConfig("(|(&(objectCategory=person)(objectClass=user)))", attrs...)
	ldapConfig.BaseDN = UserDN

	return ldapConfig.Search()
}

func SearchGroupByDN(groupDN string, attrs ...attribute) (*ldap.SearchResult, error) {

	ldapConfig := SearchConfig("(objectClass=group)", attrs...)
	ldapConfig.BaseDN = groupDN

	return ldapConfig.Search()
}

func openLogonServer() (*bufio.Scanner, *os.File, error) {
	file, err := os.Open(global.LOGON)

	if err != nil {
		logger.Error(err)
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)

	return scanner, file, err
}
