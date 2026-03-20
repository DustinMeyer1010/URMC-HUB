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
	baseDN     string
	attributes []string
	scope      int
	deref      int
	sizeLimit  int
	timeLimit  int
	typesOnly  bool
	filter     string
	control    []ldap.Control
}

// Creates a search request config with default settings:
//
//   - BaseDN:    URMC Domain
//   - Scope:     WholeSubtree
//   - Deref:     NeverDerefAliases
//   - Attribute: distinguishedName
//   - Filter:    ObjectClass=anr
//   - SizeLimit: 0
//   - TimeLimit: 0
//   - TypesOnly: false
//   - Control:   Paging (100)
func DefaultSearchConfig() LDAPSearchConfig {

	return LDAPSearchConfig{
		baseDN:     global.BASEDN,
		scope:      ldap.ScopeWholeSubtree,
		deref:      ldap.NeverDerefAliases,
		attributes: []string{"distinguishedName"},
		filter:     "(objectClass=*)",
		sizeLimit:  0,
		timeLimit:  0,
		typesOnly:  false,
		control:    []ldap.Control{ldap.NewControlPaging(50)},
	}

}

// Creates a default config but with the search only being under
// users
func UserSearchConfig() LDAPSearchConfig {
	return DefaultSearchConfig().SetFilter("(objectClass=user)")
}

// Creates a default config but with the search only being under
// computers
func ComputerSearchConfig() LDAPSearchConfig {
	return DefaultSearchConfig().SetFilter("(objectClass=computer)")
}

// Creates a default config but with the search only being under
// groups
func GroupSearchConfig() LDAPSearchConfig {
	return DefaultSearchConfig().SetFilter("(objectClass=group)")
}

func (c LDAPSearchConfig) EntryExist() (*ldap.Entry, error) {
	searchResult, err := c.Search()

	if searchResult == nil || len(searchResult.Entries) == 0 || err != nil {
		return nil, fmt.Errorf("Entry do not exists")
	}

	return searchResult.Entries[0], nil
}

// Set the Filter values for a ldap search config
func (c LDAPSearchConfig) SetFilter(filter string) LDAPSearchConfig {
	c.filter = filter
	return c
}

// Set the Scope values for a ldap search config
func (c LDAPSearchConfig) SetScope(scope int) LDAPSearchConfig {
	c.scope = scope
	return c
}

// Set the Deref values for a ldap search config
func (c LDAPSearchConfig) SetDeref(deref int) LDAPSearchConfig {
	c.deref = deref
	return c
}

// Set the SizeLimit values for a ldap search config
func (c LDAPSearchConfig) SetSizeLimit(sizeLimit int) LDAPSearchConfig {
	c.sizeLimit = sizeLimit
	return c
}

// Set the TimeLimit values for a ldap search config
func (c LDAPSearchConfig) SetTimeLimit(timeLimit int) LDAPSearchConfig {
	c.timeLimit = timeLimit
	return c
}

// Set the TypesOnly values for a ldap search config
func (c LDAPSearchConfig) SetTypesOnly(typesOnly bool) LDAPSearchConfig {
	c.typesOnly = typesOnly
	return c
}

// Set the Control values for a ldap search config
func (c LDAPSearchConfig) SetControl(control []ldap.Control) LDAPSearchConfig {
	c.control = control
	return c
}

// Set the Attributes values for a ldap search config. Finds the alias
// for the inputed attributes and outputs correct format for ldap.
func (c LDAPSearchConfig) SetAttributes(attributes []string) LDAPSearchConfig {
	c.attributes = convertAliases(attributes)
	return c
}

// Set the BaseDN values for a ldap search config
func (c LDAPSearchConfig) SetBaseDN(baseDN string) LDAPSearchConfig {
	c.baseDN = baseDN
	return c
}

// Searches ldap based on the config struct that is provided
func (c LDAPSearchConfig) Search() (*ldap.SearchResult, error) {

	LDAP_CONNECTION.CheckConnection()

	searchRequest := ldap.NewSearchRequest(
		c.baseDN,
		c.scope,
		c.deref,
		c.sizeLimit,
		c.timeLimit,
		c.typesOnly,
		c.filter,
		c.attributes,
		c.control,
	)

	return LDAP_CONNECTION.Conn.Search(searchRequest)
}

// SearchOnServer executes a defined LDAP search against a specific domain controller.
// It handles the connection lifecycle to the target server and returns the raw
// ldap.SearchResult or an error if the connection or query fails.
func (c LDAPSearchConfig) SearchOnServer(server string) (*ldap.SearchResult, error) {
	conn, cError := ConnectToServer(server)
	if cError != nil {
		return nil, cError
	}

	defer conn.Close()
	defer conn.Unbind()

	searchRequest := ldap.NewSearchRequest(
		c.baseDN,
		c.scope,
		c.deref,
		c.sizeLimit,
		c.timeLimit,
		c.typesOnly,
		c.filter,
		c.attributes,
		c.control,
	)

	return conn.Search(searchRequest)
}

type LDAPModifyConfig struct {
	DistinguishedName string
	Control           []ldap.Control
}

// NewDefaultModifyConfig creates a new LDAPModifyConfig with the target DN
// and no additional LDAP controls.
func NewDefaultModifyConfig(dn string) LDAPModifyConfig {
	return LDAPModifyConfig{
		DistinguishedName: dn,
		Control:           nil,
	}
}

// Set the control of a modify config
func (c LDAPModifyConfig) SetControl(control []ldap.Control) LDAPModifyConfig {
	c.Control = control
	return c
}

// Add appends the specified Distinguished Name (DN) to the "member" attribute
// of the LDAP object defined in the LDAPModifyConfig. It handles the full
// connection lifecycle, including connecting, binding, and unbinding.
func (c LDAPModifyConfig) Add(dn string) error {

	LDAP_CONNECTION.mutex.Lock()
	defer LDAP_CONNECTION.mutex.Unlock()

	LDAP_CONNECTION.CheckConnection()

	addRequest := ldap.NewModifyRequest(c.DistinguishedName, c.Control)
	addRequest.Add("member", []string{dn})
	ldapError := LDAP_CONNECTION.Conn.Modify(addRequest)

	if ldapError != nil {
		return ldapError
	}

	return nil

}

// Removes the specified Distinguished Name (DN) from the "member" attribute
// of the LDAP object defined in the LDAPModifyConfig. It handles the full
// connection lifecycle, including connecting, binding, and unbinding.
func (c LDAPModifyConfig) Remove(dn string) error {
	LDAP_CONNECTION.mutex.Lock()
	defer LDAP_CONNECTION.mutex.Unlock()

	removeRequest := ldap.NewModifyRequest(c.DistinguishedName, c.Control)
	removeRequest.Delete("member", []string{dn})
	ldapError := LDAP_CONNECTION.Conn.Modify(removeRequest)

	if ldapError != nil {
		return ldapError
	}

	return nil
}

// Deprecated: This is being replace with LDAPSearchConfig
func SearchAllByCategory(category, attribute, value string, attrs ...string) (*ldap.SearchResult, error) {

	filter := fmt.Sprintf("(&(objectCategory=%s)(%s=%s*))", category, attribute, value)

	ldapConfig := DefaultSearchConfig().SetFilter(filter).SetAttributes(attrs)

	return ldapConfig.Search()
}

// Deprecated: This is being replace with LDAPSearchConfig
func SearchByCategory(category, attribute, value string, attrs ...string) (*ldap.SearchResult, error) {

	filter := fmt.Sprintf("(&(objectCategory=%s)(%s=%s))", category, attribute, value)

	ldapConfig := DefaultSearchConfig().SetFilter(filter).SetAttributes(attrs)

	return ldapConfig.Search()
}

// Deprecated: This is being replace with LDAPSearchConfig
func SearchWithFilter(filter string, attrs ...string) (*ldap.SearchResult, error) {

	ldapConfig := DefaultSearchConfig().SetFilter(filter).SetAttributes(attrs)

	return ldapConfig.Search()
}

// Deprecated: This is being replace with LDAPSearchConfig
func SearchByUserDN(UserDN string, attrs ...string) (*ldap.SearchResult, error) {

	ldapConfig := DefaultSearchConfig().SetFilter("(|(&(objectCategory=person)(objectClass=user)))").SetAttributes(attrs).SetBaseDN(UserDN)

	return ldapConfig.Search()
}

// Deprecated: This is being replace with LDAPSearchConfig
func SearchGroupByDN(groupDN string, attrs ...string) (*ldap.SearchResult, error) {

	ldapConfig := DefaultSearchConfig().SetFilter("(objectClass=group)").SetAttributes(attrs).SetBaseDN(groupDN)

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
