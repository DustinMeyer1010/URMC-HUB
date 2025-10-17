package models

import "github.com/go-ldap/ldap/v3"

type GroupSimpleInfo struct {
	Name        string `json:"name"`
	OU          string `json:"ou"`
	Information string `json:"information"`
	Description string `json:"description"`
}

func ToGroupSimpleInfo(entry *ldap.Entry) GroupSimpleInfo {
	return GroupSimpleInfo{
		Name:        entry.GetAttributeValue("sAMAccountName"),
		OU:          entry.GetAttributeValue("distinguishedName"),
		Description: entry.GetAttributeValue("description"),
		Information: entry.GetAttributeValue("info"),
	}
}
