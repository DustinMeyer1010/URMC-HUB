package models

import "github.com/go-ldap/ldap/v3"

// GroupSimpleInfo represents a minimal view of an Active Directory group.
// It contains key identifying and descriptive fields that the frontend uses
// to display group search results or summaries.
type GroupSimpleInfo struct {
	Name        string `json:"name"`
	OU          string `json:"ou"`
	Information string `json:"information"`
	Description string `json:"description"`
}

// Converts a ldap entry into the a simple group info
func ToGroupSimpleInfo(entry *ldap.Entry) GroupSimpleInfo {
	return GroupSimpleInfo{
		Name:        entry.GetAttributeValue("sAMAccountName"),
		OU:          entry.GetAttributeValue("distinguishedName"),
		Description: entry.GetAttributeValue("description"),
		Information: entry.GetAttributeValue("info"),
	}
}
