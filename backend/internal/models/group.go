package models

import "github.com/go-ldap/ldap/v3"

// GroupSimpleInfo represents a minimal view of an Active Directory group.
// It contains key identifying and descriptive fields that the frontend uses
// to display group search results or summaries.
type GroupSimpleInfo struct {
	Name        string `json:"name"`
	Information string `json:"information"`
	Description string `json:"description"`
	OU          string `json:"ou"`
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

// Members to modified either to be added or removed from a group
type ModifyMembers struct {
	Members []string `json:"members"`
}

type GroupModifyRequest struct {
	UserDN []string `json:"user"`
}

// Used to show Remove and Add of group results
type GroupModifyResults struct {
	Group      string `json:"group"`
	User       string `json:"user"`
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
}
