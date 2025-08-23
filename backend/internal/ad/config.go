package ad

import "github.com/go-ldap/ldap/v3"

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
func SearchConfig(attribute []string, filter string) LDAPSearchConfig {

	return LDAPSearchConfig{
		BaseDN:    "DC=URMC-sh,DC=rochester,DC=edu",
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
