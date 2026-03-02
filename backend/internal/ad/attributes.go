package ad

import (
	"fmt"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

var attributeAliases = map[string]string{
	"CN":         "cn",
	"COMMONNAME": "cn",

	"NAME": "name",

	"USERNAME":       "sAMAccountName",
	"SAMACCOUNTNAME": "sAMAccountName",

	"DN":                "distinguishedName",
	"DISTINGUSHIEDNAME": "distinguishedName",

	"NETID": "uid",
	"UID":   "uid",

	"URID": "URID",

	"EMAIL": "mail",
	"MAIL":  "mail",

	"DEPARTMENT": "department",
	"DEPT":       "department",

	"TITLE": "title",

	"PHONE":           "telephoneNumber",
	"PHONENUMBER":     "telephoneNumber",
	"TELEPHONENUMBER": "telephoneNumber",

	"PWDLASTSET":      "pwdlastset",
	"PASSWORDLASTSET": "pwdlastset",

	"DESCRIPTION": "description",
	"DESC":        "description",

	"LOCATION":                  "physicalDeliveryOfficeName",
	"PYSICALDELIVERYOFFICENAME": "physicalDeliveryOfficeName",

	"FIRST":     "givenName",
	"FIRSTNAME": "givenName",

	"LAST":     "sn",
	"LASTNAME": "sn",

	"RELATIONSHIP":       "URRoleStatus",
	"URROLESTATUS":       "URRoleStatus",
	"RELATIONSHIPSTATUS": "URRoleStatus",
	"STATUS":             "URRoleStatus",

	"OS":              "operatingSystem",
	"OPERATINGSYSTEM": "operatingSystem",

	"INFORMATION": "information",
	"INFO":        "information",

	"BADPASSWORDCOUNT": "badPwdCount",
	"BADPWDCOUNT":      "badPwdCount",

	"BADPASSWORDTIME": "badPasswordTime",
	"BADPWDTIME":      "badPasswordTime",

	"MEMBEROF": "memberof",
}

// Pulls ldap attributes from alias, If alias does not exist
// return the orginal values
func GetAttributeAlias(attribute string) string {
	attr, ok := attributeAliases[strings.ToUpper(attribute)]

	if !ok {
		return attribute
	}

	return attr
}

// Converts alias name of ldap search to the correct
// formated names for ldap search
func convertAliases(attributes []string) []string {
	attr := []string{}
	for _, a := range attributes {
		if a == "*" {
			return []string{"*"}
		}
		attr = append(attr, GetAttributeAlias(a))
	}

	fmt.Println(attr)
	return attr
}

func createAttributeMapping(entry *ldap.Entry, attributes []string) map[string][]string {
	attrs := map[string][]string{}

	for i, a := range convertAliases(attributes) {
		if a == "*" {
			return allAttributes(entry)
		}
		attrs[attributes[i]] = entry.GetAttributeValues(a)
	}

	return attrs

}

func allAttributes(entry *ldap.Entry) map[string][]string {
	attrs := map[string][]string{}
	for _, a := range entry.Attributes {
		attrs[a.Name] = entry.GetAttributeValues(a.Name)
	}

	return attrs
}

var (
	COMMON_NAME        = "cn"
	FULL_NAME          = "name"
	USERNAME           = "sAMAccountName"
	DISTINGUISHED_NAME = "distinguishedName"
	NETID              = "uid"
	URID               = "urid"
	EMAIL              = "mail"
	DEPARMTENT         = "department"
	TITLE              = "title"
	PHONE_NUMBER       = "telephoneNumber"
	LAST_PASSWORD_SET  = "pwdlastset"
	DESCRIPTION        = "description"
	LOCATION           = "physicalDeliveryOfficeName"
	FIRST_NAME         = "givenName"
	LAST_NAME          = "sn"
	STATUS             = "URRoleStatus"
	OPERATING_SYSTEM   = "operatingSystem"
	INFORMATION        = "info"
	BAD_PASSWORD_COUNT = "badPwdCount"
	BAD_PASSWORD_TIME  = "badPasswordTime"
	MEMBER_OF          = "memberOf"
)

func memberRangeAtrribute(start, end int) string {
	return fmt.Sprintf("member;range=%d-%d", start, end)
}
