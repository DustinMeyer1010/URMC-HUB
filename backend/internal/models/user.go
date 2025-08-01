package models

import "github.com/go-ldap/ldap/v3"

type UserSimpleInfo struct {
	Type     string `json:"Type"`
	Name     string `json:"Name"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
	NetID    string `json:"NetID"`
	OU       string `json:"OU"`
}

func (u *UserSimpleInfo) FillAttributes(user *ldap.Entry) {
	u.Name = user.GetAttributeValue("Name")
	u.Username = user.GetAttributeValue("Username")
	u.Email = user.GetAttributeValue("mail")
	u.NetID = user.GetAttributeValue("urid")
	u.OU = user.GetAttributeValue("distinguishedName")
}

type UserFullInfo struct {
	Username           string `json:"Username"`
	Name               string `json:"Name"`
	Email              string `json:"Email"`
	RelationshipStatus string `json:"RelationshipStatus"`
	Department         string `json:"Department"`
	Title              string `json:"Title"`
	Phone              string `json:"Phone"`
	Location           string `json:"Location"`
	LastPasswordSet    string `json:"LastPasswordSet"`
	URID               string `json:"URID"`
	OU                 string `json:"OU"`
	NetID              string `json:"NetID"`
	Description        string `json:"Description"`
}

func (u *UserFullInfo) FillAttributes(user *ldap.Entry) {
	u.Name = user.GetAttributeValue("cn")
	u.Username = user.GetAttributeValue("Username")
	u.NetID = user.GetAttributeValue("uid")
	u.URID = user.GetAttributeValue("URID")
	u.Email = user.GetAttributeValue("mail")
	u.Phone = user.GetAttributeValue("telephoneNumber")
	u.Department = user.GetAttributeValue("department")
	u.Title = user.GetAttributeValue("title")
	u.OU = user.GetAttributeValue("distinguishedName")
	u.LastPasswordSet = user.GetAttributeValue("pwdLastSet")
	u.RelationshipStatus = user.GetAttributeValue("URRoleStatus")
	u.Location = user.GetAttributeValue("physicalDeliveryOfficeName")
	u.Description = user.GetAttributeValue("description")
}
