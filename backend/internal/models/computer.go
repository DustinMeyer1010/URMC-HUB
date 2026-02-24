package models

import "github.com/go-ldap/ldap/v3"

// Simple info for the cards of computers
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

// Page information for printers
type ComputerPageInfo struct {
	ComputerInfo ComputerSimpleInfo `json:"computer_info"`
	IsOnline     bool               `json:"is_online"`
	PingResults  string             `json:"ping_results"`
}
