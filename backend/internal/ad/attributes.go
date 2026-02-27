package ad

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

type attribute struct {
	LdapName string
}

var (
	COMMON_NAME        = attribute{LdapName: "cn"}
	FULL_NAME          = attribute{LdapName: "name"}
	USERNAME           = attribute{LdapName: "sAMAccountName"}
	DISTINGUISHED_NAME = attribute{LdapName: "distinguishedName"}
	NETID              = attribute{LdapName: "uid"}
	URID               = attribute{LdapName: "urid"}
	EMAIL              = attribute{LdapName: "mail"}
	DEPARMTENT         = attribute{LdapName: "department"}
	TITLE              = attribute{LdapName: "title"}
	PHONE_NUMBER       = attribute{LdapName: "telephoneNumber"}
	LAST_PASSWORD_SET  = attribute{LdapName: "pwdlastset"}
	DESCRIPTION        = attribute{LdapName: "description"}
	LOCATION           = attribute{LdapName: "physicalDeliveryOfficeName"}
	FIRST_NAME         = attribute{LdapName: "givenName"}
	LAST_NAME          = attribute{LdapName: "sn"}
	STATUS             = attribute{LdapName: "URRoleStatus"}
	OPERATING_SYSTEM   = attribute{LdapName: "operatingSystem"}
	INFORMATION        = attribute{LdapName: "info"}
	BAD_PASSWORD_COUNT = attribute{LdapName: "badPwdCount"}
	BAD_PASSWORD_TIME  = attribute{LdapName: "badPasswordTime"}
	MEMBER_OF          = attribute{LdapName: "memberOf"}
)

// Converts an array of attribute structs into an array of strings for ldap searching
func attributesToStrings(attrs ...attribute) []string {
	var stringAttributes = []string{}
	for _, attr := range attrs {
		stringAttributes = append(stringAttributes, attr.LdapName)
	}
	return stringAttributes

}

func createAttrToValueMapping(entry *ldap.Entry, attrs ...attribute) map[attribute]string {
	var result map[attribute]string = map[attribute]string{}
	for _, attr := range attrs {
		result[attr] = entry.GetAttributeValue(attr.LdapName)
	}
	return result
}

func memberRangeAtrribute(start, end int) attribute {
	return attribute{LdapName: (fmt.Sprintf("member;range=%d-%d", start, end))}
}
