package models

import "github.com/go-ldap/ldap/v3"

type ComputerSimpleInfo struct {
	Name            string `json:"name"`
	OU              string `json:"ou"`
	OperatingSystem string `json:"operating_system"`
}

// Converts the ldap entry into the computer simple info
func ToComputerSimpleInfo(entry *ldap.Entry) ComputerSimpleInfo {
	return ComputerSimpleInfo{
		Name:            entry.GetAttributeValue("name"),
		OU:              entry.GetAttributeValue("distinguishedName"),
		OperatingSystem: entry.GetAttributeValue("operatingSystem"),
	}
}
